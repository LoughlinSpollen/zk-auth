package usecase

import (
	"errors"
	"flag"
	"math/big"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	errormsg "zk_auth_service/pkg/domain/errors"
	model "zk_auth_service/pkg/domain/model"
	infra "zk_auth_service/pkg/infra"
	"zk_auth_service/pkg/infra/env"

	cpd "zk_auth_service/lib/zk_auth_lib"
)

var (
	fs = flag.NewFlagSet("rpc-services", flag.ExitOnError)
	q  = fs.Int64("q", env.WithDefaultInt64("ZK_CPD_Q", 10003), "q value for Chaum Pedersen")
	g  = fs.Int64("g", env.WithDefaultInt64("ZK_CPD_G", 1025), "g value for Chaum Pedersen")
	a  = fs.Int64("a", env.WithDefaultInt64("ZK_CPD_A", 1025), "a value for Chaum Pedersen")
	b  = fs.Int64("b", env.WithDefaultInt64("ZK_CPD_B", 1025), "b value for Chaum Pedersen")
)

// Onion architecture:
// This might be considered overkill for a small project, but it is a good practice to separate the layers.
// The usecase layer is the business logic layer.

// Usecases are handlers that are used to interact with the infra layer and the domain layer
// They wire up both layers through IoC.
// In this case the Chaum Pedersen verifier is (could be viewed as) a domain service.

// I use the word `usecases` because it maps easily to the business requriements.
// On projects with non technical people, it is easier to explain the usecases.

type AuthUsecase interface {
	Register(userID string, y1, y2 *big.Int) error
	AuthChallenge(userID string) (uuid.UUID, *big.Int, error)
	Authenticate(userID uuid.UUID, z *big.Int) (uuid.UUID, error)
}

type authUsecase struct {
	authRepo infra.AuthRepositoryAdapter
	verifier cpd.ChaumPedersenVerifier
	cache    infra.SessionCacheAdapter
}

func NewAuthUsecase(authRepo infra.AuthRepositoryAdapter, cache infra.SessionCacheAdapter) *authUsecase {
	return &authUsecase{
		authRepo: authRepo,
		verifier: cpd.NewChaumPedersen(big.NewInt(*g), big.NewInt(*q), big.NewInt(*a), big.NewInt(*b)),
		cache:    cache,
	}
}

func (u *authUsecase) Register(userID string, y1, y2 *big.Int) error {
	log.Trace("authUsecase Register")

	_, err := u.authRepo.Save(userID, y1, y2)
	if err != nil {
		if err.Error() == errormsg.ErrExists {
			_, err2 := u.authRepo.Update(userID, y1, y2)
			if err2 != nil {
				return err2
			}
		}
	}
	return err
}

func (u *authUsecase) AuthChallenge(userID string) (uuid.UUID, *big.Int, error) {
	log.Trace("authUsecase AuthChallenge")

	authID, err := u.authRepo.ReadAuthID(userID)
	if err != nil {
		return uuid.Nil, nil, err
	}

	s := u.verifier.Challenge()
	return authID, s, nil
}

func (u *authUsecase) Authenticate(userID uuid.UUID, z *big.Int) (uuid.UUID, error) {
	log.Trace("authUsecase Authenticate")

	y1, y2, err := u.authRepo.Read(userID)
	if err != nil {
		return uuid.Nil, err
	}
	ok := u.verifier.Verify(y1, y2, z)
	if ok {
		return uuid.Nil, errors.New(errormsg.ErrAuthFailed)
	}
	session := model.NewSession(userID)
	err = u.cache.Save(session)
	if err != nil {
		return uuid.Nil, err
	}

	return session.ID, nil
}
