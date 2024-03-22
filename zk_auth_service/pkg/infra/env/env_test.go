package env_test

import (
	"os"

	"testing"

	"zk_auth_service/pkg/infra/env"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfraEnv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "env config unit test suite")
}

var _ = Describe("infrastructure environment variables", func() {
	Describe("Reading env variables for REST service", func() {
		Context("when the AUTH_API_HTTP_PORT environment variable is not a number", func() {
			BeforeEach(func() {
				os.Setenv("AUTH_API_HTTP_PORT", "-")
			})
			AfterEach(func() {
				os.Setenv("AUTH_API_HTTP_PORT", "8080")
			})
			It("should return the default value", func() {
				result := env.WithDefaultInt("AUTH_API_HTTP_PORT", 8080)
				Expect(result).Should(Equal(8080))
			})
		})
	})

	Describe("Reading env variables for int64 value", func() {
		Context("when the int64 environment variable is valid", func() {
			BeforeEach(func() {
				os.Setenv("INT_64_VAL", "9223372036854775807")
			})
			AfterEach(func() {
				os.Unsetenv("INT_64_VAL")
			})
			It("should return the env variable value", func() {
				result := env.WithDefaultInt64("INT_64_VAL", int64(0))
				Expect(result).Should(Equal(int64(9223372036854775807)))
			})
		})
	})

	Describe("Reading env variables for int value", func() {
		Context("when the int environment variable is valid", func() {
			BeforeEach(func() {
				os.Setenv("INT_VAL", "10")
			})
			AfterEach(func() {
				os.Unsetenv("INT_VAL")
			})
			It("should return the env variable value", func() {
				result := env.WithDefaultInt("INT_VAL", 0)
				Expect(result).Should(Equal(10))
			})
		})
	})
})
