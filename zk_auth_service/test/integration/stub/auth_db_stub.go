package integration_test_stub

import (
	"errors"
	"math/big"
	errormsgs "zk_auth_service/pkg/domain/errors"

	"github.com/google/uuid"
)

type authDBStub struct {
	id     uuid.UUID
	userID string
	y1, y2 *big.Int
}

func NewDbStub() *authDBStub {
	return &authDBStub{}
}

func (stub *authDBStub) Connect() error {
	return nil
}

func (stub *authDBStub) Close() {
}

func (stub *authDBStub) Save(userID string, y1, y2 *big.Int) (uuid.UUID, error) {
	stub.id = uuid.MustParse("cd27e265-9605-4b4b-a0e5-3003ea9cc4da")
	stub.userID = userID
	stub.y1 = y1
	stub.y2 = y2
	return stub.id, nil
}

func (stub *authDBStub) Update(userID string, y1, y2 *big.Int) (uuid.UUID, error) {
	stub.y1 = y1
	stub.y2 = y2
	return stub.id, nil
}

func (stub *authDBStub) Read(authID uuid.UUID) (*big.Int, *big.Int, error) {
	return stub.y1, stub.y2, nil
}

func (stub *authDBStub) ReadAuthID(userID string) (uuid.UUID, error) {
	return stub.id, nil
}

type errDBStub struct {
}

// Error stub for database

func NewDbErrStub() *errDBStub {
	return &errDBStub{}
}

func (stub *errDBStub) Connect() error {
	return nil
}

func (stub *errDBStub) Close() {
}

func (stub *errDBStub) Save(userID string, y1, y2 *big.Int) (uuid.UUID, error) {

	return uuid.Nil, errors.New(errormsgs.ErrExists)
}

func (stub *errDBStub) Update(userID string, y1, y2 *big.Int) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (stub *errDBStub) Read(authID uuid.UUID) (*big.Int, *big.Int, error) {
	return nil, nil, nil
}

func (stub *errDBStub) ReadAuthID(userID string) (uuid.UUID, error) {
	return uuid.Nil, nil
}
