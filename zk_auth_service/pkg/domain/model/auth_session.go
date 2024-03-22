package model

import (
	"github.com/google/uuid"
)

type AuthSession struct {
	ID     uuid.UUID
	AuthID uuid.UUID
}

func NewSession(authID uuid.UUID) *AuthSession {
	return &AuthSession{
		ID:     uuid.New(),
		AuthID: authID,
	}
}
