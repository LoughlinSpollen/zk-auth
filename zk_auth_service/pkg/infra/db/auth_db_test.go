package auth_db_test

import (
	"math/big"

	infra "zk_auth_service/pkg/infra"
	database "zk_auth_service/pkg/infra/db"

	"github.com/google/uuid"

	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDatabaseService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth db unit test suite")
}

var _ = Describe("Auth Database", func() {
	var (
		db     infra.AuthRepositoryAdapter
		userID string
		authID uuid.UUID
	)

	BeforeEach(func() {
		db = database.NewAuthDB()
		err := db.Connect()
		Expect(err).ShouldNot(HaveOccurred())
	})
	AfterEach(func() {
		db.Close()
		// TODO
		// clean up database / cascade delete
		// or mock the database
	})
	Describe("create a new user", func() {
		Context("with userID and y1, y2", func() {
			BeforeEach(func() {
				userID = "user1"
			})

			It("returns auth ID without error", func() {
				var err error
				y1 := big.NewInt(1)
				y2 := big.NewInt(2)
				authID, err = db.Save(userID, y1, y2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(authID).ShouldNot(Equal(""))
			})

			It("reads saved auth attributes, without error ", func() {
				y1, y2, err := db.Read(authID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(y1).Should(Equal(big.NewInt(1)))
				Expect(y2).Should(Equal(big.NewInt(2)))
			})
		})
	})
	Describe("change the y1, y2 password ", func() {
		Context("for an existing auth user", func() {

			BeforeEach(func() {
				userID = "user2"
			})
			It("creates a new user, without error ", func() {
				var err error
				y1 := big.NewInt(1)
				y2 := big.NewInt(2)
				authID, err = db.Save(userID, y1, y2)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("updates the y1 and y2 values, without error ", func() {
				var authID2 uuid.UUID
				var err error
				y3 := big.NewInt(3)
				y4 := big.NewInt(4)
				authID2, err = db.Update(userID, y3, y4)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(authID).Should(Equal(authID2))
			})

			It("reads updated auth attributes, without error ", func() {
				y5, y6, err := db.Read(authID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(y5).Should(Equal(big.NewInt(3)))
				Expect(y6).Should(Equal(big.NewInt(4)))
			})

			It("reads the orginal auth ID, without error ", func() {
				authID2, err := db.ReadAuthID(userID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(authID).Should(Equal(authID2))
			})
		})
	})
})
