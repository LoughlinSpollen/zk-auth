package integration_test_stub

import (
	"zk_auth_service/pkg/domain/model"
)

type authCacheStub struct {
	session *model.AuthSession
}

func NewCacheStub() *authCacheStub {
	return &authCacheStub{}
}

func (stub *authCacheStub) Save(session *model.AuthSession) error {
	stub.session = session
	return nil
}

func (stub *authCacheStub) Read(sessionID string) (*model.AuthSession, error) {
	return stub.session, nil
}

func (stub *authCacheStub) Delete(sessionID string) error {
	stub.session = nil
	return nil
}
