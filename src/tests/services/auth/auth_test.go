package auth

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/sr-2020/eva-gateway/app"
	"github.com/sr-2020/eva-gateway/app/entity"
	"github.com/sr-2020/eva-gateway/tests/actors"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const (
	authLogin = "auth-service@test.com"
	authPassword = "auth-service@test.com"
)

func TestCheck(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {
		cfg := app.InitConfig()
		authService := NewAuth(cfg.Gateway + "/api/v1")

		convey.Convey("Check response", func() {
			convey.So(authService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestRegister(t *testing.T) {
	cfg := app.InitConfig()
	authService := NewAuth(cfg.Gateway + "/api/v1")

	rand.Seed((time.Now()).UnixNano())
	registerEmail := fmt.Sprintf("auth-test-%d@test.com", rand.Uint32() % 100000)

	convey.Convey("Try to register with empty creds", t, func() {
		token, statusCode, err := authService.Register(map[string]string{})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
	})

	convey.Convey("Try to register with empty password", t, func() {
		token, statusCode, err := authService.Register(map[string]string{
			"email": registerEmail,
			"password": "",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
	})

	convey.Convey("Register success", t, func() {
		token, statusCode, err := authService.Register(map[string]string{
			"email": registerEmail,
			"password": registerEmail,
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusCreated)

		oldToken := token
		convey.Convey("Login", func() {
			token, statusCode, err := authService.Auth(map[string]string{
				"email":    registerEmail,
				"password": registerEmail,
			})

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(token.Id, convey.ShouldEqual, oldToken.Id)
			convey.So(token.ApiKey, convey.ShouldNotEqual, "")
			convey.So(token.ApiKey, convey.ShouldNotEqual, oldToken.ApiKey)

			convey.Convey("Try to register with the same email", func() {
				token, statusCode, err := authService.Register(map[string]string{
					"email":    registerEmail,
					"password": registerEmail,
				})

				convey.So(err, convey.ShouldNotBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
			})

			convey.Convey("Try to change email to already exists email", func() {
				data := map[string]interface{}{
					"email": authLogin,
				}
				_, statusCode, err := authService.EditProfile(token.ApiKey, data)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
			})

			convey.Convey("Edit profile change email", func() {
				newEmail := "new-" + registerEmail
				data := map[string]interface{}{
					"email": newEmail,
				}
				user, statusCode, err := authService.EditProfile(token.ApiKey, data)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

				convey.So(user.Id, convey.ShouldEqual, token.Id)

				convey.Convey("Login with new email", func() {
					token, statusCode, err := authService.Auth(map[string]string{
						"email": newEmail,
						"password": registerEmail,
					})

					convey.So(err, convey.ShouldBeNil)
					convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

					convey.So(token.Id, convey.ShouldEqual, user.Id)
					convey.So(token.ApiKey, convey.ShouldNotEqual, "")
					convey.So(token.ApiKey, convey.ShouldNotEqual, oldToken)

					convey.Convey("Edit profile change password", func() {
						newPassword := "new-" + registerEmail
						data := map[string]interface{}{
							"password": newPassword,
						}
						user, statusCode, err := authService.EditProfile(token.ApiKey, data)

						convey.So(err, convey.ShouldBeNil)
						convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

						convey.So(user.Id, convey.ShouldEqual, token.Id)

						convey.Convey("Login with new password", func() {
							token, statusCode, err := authService.Auth(map[string]string{
								"email":    newEmail,
								"password": newPassword,
							})

							convey.So(err, convey.ShouldBeNil)
							convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

							convey.So(token.Id, convey.ShouldEqual, user.Id)
							convey.So(token.ApiKey, convey.ShouldNotEqual, "")

							convey.Convey("Read profile", func() {
								user, statusCode, err := authService.ReadProfile(token.ApiKey)

								convey.So(err, convey.ShouldBeNil)
								convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

								convey.So(user.Id, convey.ShouldEqual, token.Id)
							})
						})
					})
				})
			})
		})

		admin := actors.NewAdmin(cfg.ApiKey, cfg.Gateway + "/api/v1")

		convey.So(admin.DeleteUser(oldToken.Id), convey.ShouldBeNil)
	})
}

func TestProfile(t *testing.T) {
	cfg := app.InitConfig()
	authService := NewAuth(cfg.Gateway + "/api/v1")

	convey.Convey("Register or login with valid creds", t, func() {
		token, statusCode, err := authService.Register(map[string]string{
			"email": authLogin,
			"password": authPassword,
		})

		if statusCode != http.StatusCreated {
			token, statusCode, err = authService.Auth(map[string]string{
				"email": authLogin,
				"password": authPassword,
			})

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)
		}

		convey.So(token.Id, convey.ShouldNotEqual, 0)
		convey.So(token.ApiKey, convey.ShouldNotEqual, "")

		oldToken := token.ApiKey
		oldUserId := token.Id
		convey.Convey("Read profile", func() {
			user, statusCode, err := authService.ReadProfile(token.ApiKey)

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(user.Id, convey.ShouldEqual, oldUserId)
			convey.So(user.CreatedAt, convey.ShouldNotEqual, "")
			convey.So(user.UpdatedAt, convey.ShouldNotEqual, "")
		})

		convey.Convey("Try to edit profile set admin status", func() {
			data := map[string]interface{}{
				"admin": true,
			}
			user, statusCode, err := authService.EditProfile(token.ApiKey, data)

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(user.Id, convey.ShouldEqual, user.Id)
			convey.So(user.Admin, convey.ShouldEqual, false)

			convey.Convey("Read profile for admin", func() {
				user, statusCode, err := authService.ReadProfile(token.ApiKey)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

				convey.So(user.Id, convey.ShouldEqual, user.Id)
				convey.So(user.Admin, convey.ShouldEqual, false)
			})
		})

		convey.Convey("Try to edit profile set id", func() {
			newId := 100000
			data := map[string]interface{}{
				"id": newId,
			}
			user, statusCode, err := authService.EditProfile(token.ApiKey, data)

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(user.Id, convey.ShouldNotEqual, newId)

			convey.Convey("Read profile for id", func() {
				user, statusCode, err := authService.ReadProfile(token.ApiKey)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

				convey.So(user.Id, convey.ShouldNotEqual, newId)
			})
		})

		convey.Convey("One more time login", func() {
			token, statusCode, err := authService.Auth(map[string]string{
				"email": authLogin,
				"password": authPassword,
			})

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(token.Id, convey.ShouldEqual, oldUserId)
			convey.So(token.ApiKey, convey.ShouldNotEqual, "")
			convey.So(token.ApiKey, convey.ShouldNotEqual, oldToken)

			convey.Convey("Try to read profile with old token", func() {
				user, statusCode, err := authService.ReadProfile(oldToken)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)

				convey.So(user.Id, convey.ShouldEqual, 0)
			})

			convey.Convey("Read profile with new token", func() {
				user, statusCode, err := authService.ReadProfile(token.ApiKey)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

				convey.So(user.Id, convey.ShouldEqual, oldUserId)
				convey.So(user.CreatedAt, convey.ShouldNotEqual, "")
				convey.So(user.UpdatedAt, convey.ShouldNotEqual, "")

				convey.Convey("Edit profile set name: New Name", func() {
					newName := "newName"
					data := map[string]interface{}{
						"name": newName,
					}
					user, statusCode, err := authService.EditProfile(token.ApiKey, data)

					convey.So(err, convey.ShouldBeNil)
					convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

					convey.So(user.Id, convey.ShouldEqual, oldUserId)
					convey.So(user.Name, convey.ShouldEqual, data["name"])

					convey.Convey("Read profile for new name", func() {
						user, statusCode, err := authService.ReadProfile(token.ApiKey)

						convey.So(err, convey.ShouldBeNil)
						convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

						convey.So(user.Id, convey.ShouldEqual, oldUserId)
						convey.So(user.Name, convey.ShouldEqual, data["name"])

						convey.Convey("Edit profile set status: new_status", func() {
							data := map[string]interface{}{
								"status": "new_status",
							}
							user, statusCode, err := authService.EditProfile(token.ApiKey, data)

							convey.So(err, convey.ShouldBeNil)
							convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

							convey.So(user.Id, convey.ShouldEqual, oldUserId)
							convey.So(user.Name, convey.ShouldEqual, newName)
							convey.So(user.Status, convey.ShouldEqual, data["status"])

							convey.Convey("Read profile for new status and new name", func() {
								user, statusCode, err := authService.ReadProfile(token.ApiKey)

								convey.So(err, convey.ShouldBeNil)
								convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

								convey.So(user.Id, convey.ShouldEqual, oldUserId)
								convey.So(user.Name, convey.ShouldEqual, newName)
								convey.So(user.Status, convey.ShouldEqual, data["status"])
							})
						})
					})
				})
			})
		})
	})
}

func TestLogin(t *testing.T) {
	cfg := app.InitConfig()
	authService := NewAuth(cfg.Gateway + "/api/v1")

	convey.Convey("Try to login with empty creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
	})

	convey.Convey("Try to login for not exists account", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email": "auth-wrong@test.com",
			"password": "1234",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)
		convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
	})

	convey.Convey("Try to login with wrong creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email": authLogin,
			"password": "wrong-pass",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)
		convey.So(token, convey.ShouldResemble, entity.AuthUserToken{})
	})

	convey.Convey("Login with valid creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email": authLogin,
			"password": authPassword,
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

		convey.So(token.Id, convey.ShouldNotEqual, 0)
		convey.So(token.ApiKey, convey.ShouldNotEqual, "")

		oldToken := token
		convey.Convey("One more time login", func() {
			token, statusCode, err := authService.Auth(map[string]string{
				"email":    authLogin,
				"password": authPassword,
			})

			convey.So(err, convey.ShouldBeNil)
			convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

			convey.So(token.Id, convey.ShouldEqual, oldToken.Id)
			convey.So(token.ApiKey, convey.ShouldNotEqual, "")
			convey.So(token.ApiKey, convey.ShouldNotEqual, oldToken.ApiKey)

			convey.Convey("Try to read profile by old token", func() {
				user, statusCode, err := authService.ReadProfile(oldToken.ApiKey)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)

				convey.So(user.Id, convey.ShouldEqual, 0)
			})

			convey.Convey("Read profile by new token", func() {
				user, statusCode, err := authService.ReadProfile(token.ApiKey)

				convey.So(err, convey.ShouldBeNil)
				convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

				convey.So(user.Id, convey.ShouldEqual, token.Id)
			})
		})
	})
}
