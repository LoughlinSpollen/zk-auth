package rpc

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"net"
	pb "zk_auth_service/build/protos/zk_auth"
	errors "zk_auth_service/pkg/domain/errors"
	"zk_auth_service/pkg/infra/env"
	"zk_auth_service/pkg/usecase"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	fs       = flag.NewFlagSet("rpc", flag.ExitOnError)
	grpcPort = fs.Int("service-port", env.WithDefaultInt("SERVICE_PORT", 1025), "grpc port")
)

type AuthService interface {
	Connect()
	Close()
	Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error)
	CreateAuthenticationChallenge(ctx context.Context, in *pb.AuthenticationChallengeRequest) (*pb.AuthenticationChallengeResponse, error)
	VerifyAuthentication(ctx context.Context, in *pb.AuthenticationAnswerRequest) (*pb.AuthenticationAnswerResponse, error)
}

type zkAuthService struct {
	pb.UnimplementedAuthServer
	usecase usecase.AuthUsecase
	svr     *grpc.Server
}

func NewAuthService(usecase usecase.AuthUsecase) *zkAuthService {
	authService := &zkAuthService{
		usecase: usecase,
		svr:     grpc.NewServer(),
	}
	return authService
}

func (s zkAuthService) Connect() {
	log.Trace("zkAuthService Connect")

	pb.RegisterAuthServer(s.svr, s)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("gRPC failed to connect to port %d: %v", *grpcPort, err)
	}
	log.Fatal(s.svr.Serve(conn))
}

func (s zkAuthService) Close() {
	log.Trace("zkAuthService Close")

	s.svr.GracefulStop()
}

func (s zkAuthService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Trace("zkAuthService Register")

	userID := in.GetUser()
	y1 := in.GetY1()
	y2 := in.GetY2()

	err := s.usecase.Register(userID, big.NewInt(y1), big.NewInt(y2))
	if err != nil {
		if err.Error() == errors.ErrExists {
			return nil, status.Error(codes.AlreadyExists, "User has already been registered")
		}
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (s zkAuthService) CreateAuthenticationChallenge(ctx context.Context, in *pb.AuthenticationChallengeRequest) (*pb.AuthenticationChallengeResponse, error) {
	log.Trace("zkAuthService CreateAuthenticationChallenge")

	userID := in.GetUser()
	// Not needed - but in the gRPC definition
	// r1 := in.GetR1()
	// r2 := in.GetR2()
	authID, c, err := s.usecase.AuthChallenge(userID)
	if err != nil {
		return nil, err
	}
	return &pb.AuthenticationChallengeResponse{AuthId: authID.String(), C: c.Int64()}, nil
}

func (s zkAuthService) VerifyAuthentication(ctx context.Context, in *pb.AuthenticationAnswerRequest) (*pb.AuthenticationAnswerResponse, error) {
	log.Trace("zkAuthService VerifyAuthentication")

	authID := in.GetAuthId()
	authUID := uuid.MustParse(authID)
	sec := in.GetS()
	sessionID, err := s.usecase.Authenticate(authUID, big.NewInt(sec))
	if err != nil {
		if err.Error() == errors.ErrAuthFailed {
			return nil, status.Error(codes.PermissionDenied, "Authentication failed")
		}
		return nil, err
	}
	return &pb.AuthenticationAnswerResponse{
		SessionId: sessionID.String(),
	}, nil
}
