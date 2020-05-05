package auth

import (
	"github.com/sr-2020/eva-gateway/app/entity"
)

type AuthService interface {
	Check() bool
	Auth(map[string]string) (entity.AuthUserToken, int, error)
	Register(map[string]string) (entity.AuthUserToken, int, error)
	ReadProfile(string) (entity.ProfileUser, int, error)
	EditProfile(string, map[string]interface{}) (entity.ProfileUser, int, error)
	Delete(int) error
}
