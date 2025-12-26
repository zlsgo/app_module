package model

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidCryptedID 无效的加密 ID
	ErrInvalidCryptedID = errors.New("invalid encrypted id")
	// ErrCrypterNotSet 加密器未配置
	ErrCrypterNotSet = errors.New("id crypter not configured")
	// ErrInvalidKeyLength 无效的密钥长度
	ErrInvalidKeyLength = errors.New("key must be 16, 24, or 32 bytes")
	// ErrAuthFailed 认证失败
	ErrAuthFailed = errors.New("authentication failed")
	// ErrInvalidEntity 无效实体
	ErrInvalidEntity = errors.New("invalid entity")
	// ErrInvalidData 无效数据
	ErrInvalidData = errors.New("invalid data")
	// ErrInvalidTableName 无效表名
	ErrInvalidTableName = errors.New("invalid table name")
	// ErrInvalidPrimaryKey 无效主键
	ErrInvalidPrimaryKey = errors.New("invalid primary key")
	// ErrTooManyRows 行数过多
	ErrTooManyRows = errors.New("too many rows")
	// ErrOptimisticLock 乐观锁冲突
	ErrOptimisticLock = errors.New("optimistic lock conflict")
	// ErrSoftDeleteNotSupported 不支持软删除
	ErrSoftDeleteNotSupported = errors.New("soft delete not supported")
	// ErrHookCancelled 钩子取消操作
	ErrHookCancelled = errors.New("operation cancelled by hook")
)

// ModelError 模型错误
type ModelError struct {
	Op      string
	Table   string
	Err     error
	Message string
}

// Error 返回模型错误信息
func (e *ModelError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("model: %s on %s: %s (%v)", e.Op, e.Table, e.Message, e.Err)
	}
	return fmt.Sprintf("model: %s on %s: %v", e.Op, e.Table, e.Err)
}

// Unwrap 返回原始错误
func (e *ModelError) Unwrap() error {
	return e.Err
}

// NewModelError 创建新的模型错误
func NewModelError(op, table string, err error, message ...string) *ModelError {
	me := &ModelError{Op: op, Table: table, Err: err}
	if len(message) > 0 {
		me.Message = message[0]
	}
	return me
}

// QueryError 查询错误
type QueryError struct {
	SQL     string
	Args    []any
	Err     error
	Message string
}

// Error 返回查询错误信息
func (e *QueryError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("query error: %s (%v)", e.Message, e.Err)
	}
	return fmt.Sprintf("query error: %v", e.Err)
}

// Unwrap 返回原始错误
func (e *QueryError) Unwrap() error {
	return e.Err
}

// NewQueryError 创建新的查询错误
func NewQueryError(sql string, args []any, err error, message ...string) *QueryError {
	qe := &QueryError{SQL: sql, Args: args, Err: err}
	if len(message) > 0 {
		qe.Message = message[0]
	}
	return qe
}

// HookError 钩子错误
type HookError struct {
	Event string
	Err   error
}

// Error 返回钩子错误信息
func (e *HookError) Error() string {
	return fmt.Sprintf("hook %s failed: %v", e.Event, e.Err)
}

// Unwrap 返回原始错误
func (e *HookError) Unwrap() error {
	return e.Err
}

// NewHookError 创建新的钩子错误
func NewHookError(event string, err error) *HookError {
	return &HookError{Event: event, Err: err}
}
