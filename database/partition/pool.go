// Package partition 提供按键懒打开资源的通用池。
//
// 打开方式与关闭动作由调用方以回调提供，池只负责并发单飞、数量上限
// 和生命周期。
package partition

import (
	"errors"
	"fmt"
	"slices"
	"sync"
)

var (
	// ErrClosed 表示池已经关闭。
	ErrClosed = errors.New("partitions closed")
	// ErrLimit 表示已打开的资源数量达到配置上限。
	ErrLimit = errors.New("partition limit reached")
)

// Config 定义池的打开与关闭策略。
type Config[T any] struct {
	// Base 是空键返回的默认资源，池不会关闭它。
	Base T
	// Open 打开非空键对应的资源。同一个键同时只打开一次，失败后可以重试。
	Open func(key string) (T, error)
	// Close 关闭池自己打开的资源；为 nil 表示资源由调用方管理。
	Close func(T) error
	// Max 非空键的资源数量上限，0 或负数表示不限制。
	Max int
}

// Pool 按键提供懒打开的资源，并负责并发单飞、数量上限和生命周期。
//
// 零值不可用，必须通过 NewPool 构造。
type Pool[T any] struct {
	// mu 用标准库 RWMutex：分区 map 的读写必须严格互斥。
	// zsync.RBMutex 的读偏置快路径在 -race 下能与写重叠，不能保护 map。
	mu    sync.RWMutex
	base  T
	open  func(string) (T, error)
	close func(T) error
	max   int

	// opened 是已成功打开的非空键数量。打开期间记在 reserved 上，
	// 否则两个首次访问同时通过上限检查会超发。打开失败两者都不计。
	opened   int
	reserved int

	entries   map[string]*entry[T]
	own       []T
	closed    bool
	opening   sync.WaitGroup
	closeDone chan struct{}
	closeErr  error
}

// entry 让并发首访对同一个键共享一次打开。
//
// 用 once 而不是持锁打开：打开 SQLite 文件要 Ping，持锁会把其他键的
// 首访一起阻塞，而它们本来互不相干。
type entry[T any] struct {
	once  sync.Once
	value T
	err   error
	// ready 区分“正在打开”和“已经完成”。Keys 只返回 ready 的项，
	// 避免把半初始化的资源暴露给枚举操作。
	ready bool
}

// NewPool 创建资源池。
func NewPool[T any](c Config[T]) (*Pool[T], error) {
	if c.Open == nil {
		return nil, errors.New("partition pool: open callback is nil")
	}
	return &Pool[T]{
		base:      c.Base,
		open:      c.Open,
		close:     c.Close,
		max:       c.Max,
		entries:   make(map[string]*entry[T]),
		closeDone: make(chan struct{}),
	}, nil
}

// Get 返回键对应的资源。空键返回默认资源。
//
// 首次访问会调用 Open，并发首访只会调用一次；同一个键反复调用返回
// 同一个资源。
func (p *Pool[T]) Get(key string) (T, error) {
	var zero T

	if key == "" {
		if p.isClosed() {
			return zero, ErrClosed
		}
		return p.base, nil
	}

	p.mu.RLock()
	closed := p.closed
	item := p.entries[key]
	p.mu.RUnlock()
	if closed {
		return zero, ErrClosed
	}

	if item == nil {
		p.mu.Lock()
		if p.closed {
			p.mu.Unlock()
			return zero, ErrClosed
		}
		if item = p.entries[key]; item == nil {
			if p.max > 0 && p.opened+p.reserved >= p.max {
				p.mu.Unlock()
				return zero, ErrLimit
			}
			item = &entry[T]{}
			p.entries[key] = item
			p.reserved++
			// Add 与 closed 检查在同一把锁下，Close 开始等待后不会再 Add。
			p.opening.Add(1)
		}
		p.mu.Unlock()
	}

	item.once.Do(func() {
		defer p.opening.Done()
		value, err := p.openSafe(key)
		// 持锁写 item：Keys 与并发的 Get 都在锁内读这些字段，
		// 否则首发访问期间读 item.value 是 data race。
		p.mu.Lock()
		item.value, item.err, item.ready = value, err, true
		p.reserved--
		if err == nil {
			p.opened++
			if p.close != nil {
				p.own = append(p.own, value)
			}
		} else {
			// 本轮并发调用共享失败结果，后续调用可创建新 entry 重试。
			delete(p.entries, key)
		}
		p.mu.Unlock()
	})

	p.mu.RLock()
	value, err := item.value, item.err
	closed = p.closed
	p.mu.RUnlock()
	if closed {
		return zero, ErrClosed
	}
	return value, err
}

// openSafe 把打开回调里的 panic 转成错误。
//
// 打开发生在进程内共享的池上，一次回调 panic 不应该让 entry 停在
// 「已标记 done、但没有 value/err」的状态，否则后续 Get 会拿到零值
// 和 nil error。转成错误后走正常的失败清理与重试路径。
func (p *Pool[T]) openSafe(key string) (value T, err error) {
	defer func() {
		if r := recover(); r != nil {
			var zero T
			value = zero
			err = fmt.Errorf("打开分区[%s]时 panic: %v", key, r)
		}
	}()
	return p.open(key)
}

// Base 返回默认资源：配置里的原资源。
func (p *Pool[T]) Base() T {
	return p.base
}

// Keys 返回池已打开成功的键（含默认空键），按字典序升序。
// 正在打开或打开失败的键不会出现。
func (p *Pool[T]) Keys() []string {
	p.mu.RLock()
	keys := make([]string, 0, len(p.entries)+1)
	keys = append(keys, "")
	for key, item := range p.entries {
		if !item.ready || item.err != nil {
			continue
		}
		keys = append(keys, key)
	}
	p.mu.RUnlock()

	slices.Sort(keys)
	return keys
}

// Close 拒绝新访问，等待在途打开并关闭池自己打开的资源。
// 并发或重复调用会等待同一次关闭完成；默认资源由调用方负责。
func (p *Pool[T]) Close() error {
	p.mu.Lock()
	if p.closed {
		done := p.closeDone
		p.mu.Unlock()
		<-done
		p.mu.RLock()
		err := p.closeErr
		p.mu.RUnlock()
		return err
	}
	p.closed = true
	p.mu.Unlock()

	// 打开不持 mu，先等在途打开发布所有权，再取待关闭资源。
	p.opening.Wait()
	p.mu.Lock()
	own := p.own
	p.own = nil
	p.mu.Unlock()

	var firstErr error
	for _, value := range own {
		if err := p.close(value); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	p.mu.Lock()
	p.closeErr = firstErr
	close(p.closeDone)
	p.mu.Unlock()
	return firstErr
}

func (p *Pool[T]) isClosed() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.closed
}
