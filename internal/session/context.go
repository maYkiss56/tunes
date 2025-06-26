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

//
// type sessionKey struct{}
//
// // WithSession добавляет сессию в контекст.
// // Используется в middleware аутентификации.
// func WithSession(ctx context.Context, s Session) context.Context {
// 	return context.WithValue(ctx, sessionKey{}, s)
// }
//
// // FromContext возвращает сессию из контекста и флаг её наличия.
// // Используется в обработчиках для безопасного доступа к данным пользователя.
// func FromContext(ctx context.Context) (*Session, bool) {
// 	s, ok := ctx.Value(sessionKey{}).(Session)
// 	if !ok {
// 		return nil, false
// 	}
// 	return &s, true
// }
//
// // MustFromContext возвращает сессию из контекста или паникует, если её нет.
// // Используется в местах, где сессия обязана быть (например, после AuthMiddleware).
// func MustFromContext(ctx context.Context) *Session {
// 	s, ok := FromContext(ctx)
// 	if !ok {
// 		panic("session not found in context")
// 	}
// 	return s
// }
