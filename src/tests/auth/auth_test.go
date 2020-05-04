package auth

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

const authHost = "http://gateway.evarun.ru/api/v1"

func TestCheck(t *testing.T) {
	convey.Convey("Go to check auth service", t, func() {
		authService := NewAuth(authHost)

		convey.Convey("Check response", func() {
			convey.So(authService.Check(), convey.ShouldEqual, true)
		})
	})
}

func TestRegister(t *testing.T) {
	authService := NewAuth(authHost)

	convey.Convey("Try to register with empty creds", t, func() {
		token, statusCode, err := authService.Register(map[string]string{})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, Token{})
	})

	convey.Convey("Try to register with wrong creds", t, func() {
		token, statusCode, err := authService.Register(map[string]string{
			"email": "test@tat.ru",
			"password": "",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, Token{})
	})

	convey.Convey("Register or login with valid creds", t, func() {
		token, statusCode, err := authService.Register(map[string]string{
			"email": "test@tat.ru",
			"password": "1234",
		})

		if statusCode != http.StatusCreated {
			token, statusCode, err = authService.Auth(map[string]string{
				"email": "test@tat.ru",
				"password": "1234",
			})
		}

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

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
				"email": "test@tat.ru",
				"password": "1234",
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
	authService := NewAuth(authHost)

	convey.Convey("Try to login with empty creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusBadRequest)
		convey.So(token, convey.ShouldResemble, Token{})
	})

	convey.Convey("Try to login with wrong creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email": "wrong@tat.ru",
			"password": "1234",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusUnauthorized)
		convey.So(token, convey.ShouldResemble, Token{})
	})

	convey.Convey("Login with valid creds", t, func() {
		token, statusCode, err := authService.Auth(map[string]string{
			"email": "test@tat.ru",
			"password": "1234",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(statusCode, convey.ShouldEqual, http.StatusOK)

		convey.So(token.Id, convey.ShouldNotEqual, 0)
		convey.So(token.ApiKey, convey.ShouldNotEqual, "")
	})
}
