package account

import (
	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
)

var (
	sessionHub = zarray.NewHashMap[string, *session]()
)

type session struct {
	sessions *zarray.Maper[int64, *znet.SSE]
}

func (m *session) addSession(sse *znet.SSE) int64 {
	id := zstring.UUID()
	m.sessions.Set(id, sse)
	return id
}

func (m *session) removeSession(id int64) {
	m.sessions.Delete(id)
}

func (m *Module) newSession(c *znet.Context) (sse *znet.SSE, remove func(), err error) {
	uid := Request.UID(c)
	if uid == "" {
		return nil, nil, zerror.InvalidInput.Text("用户未登录")
	}

	session, _, _ := sessionHub.ProvideGet(uid, func() (*session, bool) {
		return &session{
			sessions: zarray.NewHashMap[int64, *znet.SSE](),
		}, true
	})

	sse = znet.NewSSE(c, func(lastID string, opts *znet.SSEOption) {
		if m.Options.SSE.HeartbeatsTime != 0 {
			opts.HeartbeatsTime = m.Options.SSE.HeartbeatsTime
		}
		if m.Options.SSE.RetryTime != 0 {
			opts.RetryTime = m.Options.SSE.RetryTime
		}
	})

	id := session.addSession(sse)

	if m.Options.SSEReconnect != nil && sse.LastEventID() != "" {
		go m.Options.SSEReconnect(uid, sse.LastEventID())
	}

	return sse, func() {
		session.removeSession(id)
		if session.sessions.Len() == 0 {
			sessionHub.Delete(uid)
		}
	}, nil
}

func SendRealtime(uid string, data string, event ...string) bool {
	if session, ok := sessionHub.Get(uid); ok {
		if session.sessions.Len() == 0 {
			return false
		}
		id := ztime.Now()
		session.sessions.ForEach(func(k int64, v *znet.SSE) bool {
			_ = v.Send(id, data, event...)
			return true
		})
		return true
	}
	return false
}
