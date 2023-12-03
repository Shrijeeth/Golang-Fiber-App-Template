package authentication

import (
	"fmt"

	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/models"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/app/src/services/authentication"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/configs"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils"
	"github.com/Shrijeeth/Golang-Fiber-App-Template/pkg/utils/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Authentication Service", func() {
	Context("Get User By Email - Success Cases", func() {
		It("Success case with valid email", func ()  {
			err := configs.DbClient.Transaction(func(tx *gorm.DB) error {
				sampleUser := models.User{
					Username: "test",
					Email: "test@gmail.com",
					PasswordHash: "12345",
				}
				result := tx.Create(&sampleUser)
				if (result.Error != nil) || (result.RowsAffected == utils.IntZero) {
					tx.Rollback()
					return fmt.Errorf("error adding test user: %s", result.Error)
				}

				return nil
			})
			Expect(err).To(BeNil())

			user, err := authentication.GetUserByEmail("test@gmail.com")

			Expect(err).To(BeNil())
			Expect(user.Username).To(Equal("test"))
			Expect(user.Email).To(Equal("test@gmail.com"))
			Expect(user.PasswordHash).To(Equal("12345"))
			Expect(user.UserStatus).To(Equal(types.ActiveUser))
			Expect(user.UserRole).To(Equal(types.User))
			Expect(user.AuthenticationType).To(Equal(types.NormalAuthentication))
		})
	})

	Context("Get User By Email - Failure Case", func() {
		It("User not found for given email id", func ()  {
			user, err := authentication.GetUserByEmail("test@gmail.com")

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("record not found"))
			Expect(user).To(Equal(&models.User{}))
		})
	})
})