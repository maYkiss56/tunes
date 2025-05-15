package session

import (
	"context"
)

type sessionKey struct{}

func WithSession(ctx context.Context, s Session) context.Context {
	return context.WithValue(ctx, sessionKey{}, s)
}

func FromContext(ctx context.Context) *Session {
	s, ok := ctx.Value(sessionKey{}).(Session)
	if !ok {
		return nil
	}
	return &s
}
