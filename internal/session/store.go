package session

import "sync"

// TODO: redis
// global storage with mutex
var (
	sessionStore = make(map[string]Session)
	mu           sync.RWMutex
)

// SaveSession сохраняет сессию в памяти
func SaveSession(s Session) {
	mu.Lock()
	defer mu.Unlock()
	sessionStore[s.ID] = s
}

// GetSession возвращает сессию по ID
func GetSession(id string) (Session, bool) {
	mu.RLock()
	defer mu.RUnlock()
	s, ok := sessionStore[id]
	return s, ok
}

// DeleteSession удаляет сессию
func DeleteSession(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessionStore, id)
}
