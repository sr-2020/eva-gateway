package auth

import "github.com/sr-2020/eva-gateway/app/entity"

type AuthService interface {
	Check() bool
	Auth(map[string]string) (Token, int, error)
	Register(map[string]string) (Token, int, error)
	ReadProfile(string) (entity.ProfileUser, int, error)
	EditProfile(string, map[string]interface{}) (entity.ProfileUser, int, error)
}
