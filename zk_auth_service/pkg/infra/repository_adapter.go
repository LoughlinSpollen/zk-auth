package adapter

import (
	"math/big"
	"zk_auth_service/pkg/domain/model"

	"github.com/google/uuid"
)

type AuthRepositoryAdapter interface {
	// connect to the database
	Connect() error
	// close the database connection
	Close()
	// saves the y1 and y2 values for a given userID and returns the authID
	Save(userID string, y1, y2 *big.Int) (uuid.UUID, error)
	// updates the y1 and y2 values for a given userID and returns the existing authID
	Update(userID string, y1, y2 *big.Int) (uuid.UUID, error)
	// given a authID, return the y1 and y2 values
	Read(authID uuid.UUID) (*big.Int, *big.Int, error)
	// given a userID, return the authID
	ReadAuthID(userID string) (uuid.UUID, error)
}

type SessionCacheAdapter interface {
	// save the session
	Save(session *model.AuthSession) error
	// read the session
	Read(sessionID string) (*model.AuthSession, error)
	// delete the session
	Delete(sessionID string) error
}
