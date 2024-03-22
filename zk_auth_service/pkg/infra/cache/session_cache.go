package cache

import (
	"sync"
	"zk_auth_service/pkg/domain/model"
)

var (
	sessions      = make(map[string]*model.AuthSession)
	sessionsMutex = sync.RWMutex{}
)

type sessionCache struct{}

func NewSessionCache() *sessionCache {
	return &sessionCache{}
}

func (c *sessionCache) Save(session *model.AuthSession) error {
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()

	sessions[session.ID.String()] = session
	return nil
}

func (c *sessionCache) Read(sessionID string) (*model.AuthSession, error) {
	sessionsMutex.RLock()
	defer sessionsMutex.RUnlock()

	session, ok := sessions[sessionID]
	if !ok {
		return nil, nil
	}
	return session, nil
}

func (c *sessionCache) Delete(sessionID string) error {
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()

	delete(sessions, sessionID)
	return nil
}
