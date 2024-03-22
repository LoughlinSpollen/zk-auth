package integration_test

import (
	"context"
	"zk_auth_service/pkg/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"testing"

	pb "zk_auth_service/build/protos/zk_auth"
	infra "zk_auth_service/pkg/infra"
	rpc "zk_auth_service/pkg/infra/network/rpc"
	stub "zk_auth_service/test/integration/stub"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZK Auth server integration test suite")
}

var authService rpc.AuthService

var _ = Describe("Auth server entry points", func() {
	var (
		db          infra.AuthRepositoryAdapter
		cache       infra.SessionCacheAdapter
		authUsecase usecase.AuthUsecase
		conn        *grpc.ClientConn
		client      pb.AuthClient
		authID      string
	)

	Describe("auth registeration gRPC end-point", func() {
		BeforeEach(func() {
			db = stub.NewDbStub()
			cache = stub.NewCacheStub()
			authUsecase = usecase.NewAuthUsecase(db, cache)
			authService = rpc.NewAuthService(authUsecase)

			address := ":1025"
			var err error
			conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			Expect(err).ShouldNot(HaveOccurred())
			client = pb.NewAuthClient(conn)
		})

		AfterEach(func() {
			db.Close()
			conn.Close()
		})

		Context("register a new user", func() {
			It("successfully registered and returned", func() {
				ctx := context.Background()
				_, err := client.Register(ctx, &pb.RegisterRequest{
					User: "username",
				})
				Expect(err).ShouldNot(HaveOccurred())
				e, ok := status.FromError(err)
				Expect(ok).Should(BeTrue())
				Expect(e.Code()).Should(Equal(codes.OK))
			})
		})

		Context("register an existing user", func() {
			BeforeEach(func() {
				db = stub.NewDbErrStub()
				cache := stub.NewCacheStub()
				authUsecase = usecase.NewAuthUsecase(db, cache)
				authService = rpc.NewAuthService(authUsecase)

				address := ":1025"
				var err error
				conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
				Expect(err).ShouldNot(HaveOccurred())
				client = pb.NewAuthClient(conn)
			})

			AfterEach(func() {
				db.Close()
				conn.Close()
			})

			It("fails registations and returns appropriate message", func() {
				ctx := context.Background()
				_, err := client.Register(ctx, &pb.RegisterRequest{
					User: "username",
				})
				Expect(err).Should(HaveOccurred())
				e, ok := status.FromError(err)
				Expect(ok).Should(BeTrue())
				Expect(e.Code()).Should(Equal(codes.AlreadyExists))
			})
		})
	})

	Describe("the Authentication Challenge end-points", func() {
		Context("request a challenge", func() {
			BeforeEach(func() {

				db = stub.NewDbStub()
				cache := stub.NewCacheStub()
				authUsecase = usecase.NewAuthUsecase(db, cache)
				authService = rpc.NewAuthService(authUsecase)

				address := ":1025"
				var err error
				conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
				Expect(err).ShouldNot(HaveOccurred())
				client = pb.NewAuthClient(conn)
			})

			AfterEach(func() {
				db.Close()
				conn.Close()
			})

			It("successfully requested a challenge", func() {
				ctx := context.Background()
				response, err := client.CreateAuthenticationChallenge(ctx, &pb.AuthenticationChallengeRequest{
					User: "username",
				})
				Expect(err).ShouldNot(HaveOccurred())
				e, ok := status.FromError(err)
				Expect(ok).Should(BeTrue())
				Expect(e.Code()).Should(Equal(codes.OK))

				Expect(response).ShouldNot(BeNil())
				Expect(response.GetC()).ShouldNot(BeNil())
				Expect(response.GetAuthId()).ShouldNot(BeNil())

				authID := response.GetAuthId()
				_, err2 := uuid.Parse(authID)
				Expect(err2).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("the Authentication Verification end-points", func() {
		Context("verify a user", func() {
			BeforeEach(func() {

				db = stub.NewDbStub()
				cache := stub.NewCacheStub()
				authUsecase = usecase.NewAuthUsecase(db, cache)
				authService = rpc.NewAuthService(authUsecase)

				address := ":1025"
				var err error
				conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
				Expect(err).ShouldNot(HaveOccurred())
				client = pb.NewAuthClient(conn)

				ctx := context.Background()
				response, err := client.CreateAuthenticationChallenge(ctx, &pb.AuthenticationChallengeRequest{
					User: "username",
				})
				Expect(err).ShouldNot(HaveOccurred())
				authID = response.GetAuthId()
			})

			AfterEach(func() {
				db.Close()
				conn.Close()
			})

			It("successfully verified a user", func() {
				ctx := context.Background()
				response, err := client.VerifyAuthentication(ctx, &pb.AuthenticationAnswerRequest{
					AuthId: authID,
					S:      1234567890,
				})
				Expect(err).ShouldNot(HaveOccurred())
				e, ok := status.FromError(err)
				Expect(ok).Should(BeTrue())
				Expect(e.Code()).Should(Equal(codes.OK))

				Expect(response).ShouldNot(BeNil())
				session := response.GetSessionId()
				Expect(session).ShouldNot(BeNil())

				_, err2 := uuid.Parse(session)
				Expect(err2).ShouldNot(HaveOccurred())
			})
		})
	})
})
