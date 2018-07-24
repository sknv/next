package globals

import (
	"github.com/go-chi/jwtauth"

	"github.com/sknv/next/app/core/cfg"
)

const (
	alg = "HS256"
)

var (
	jwtAuth *jwtauth.JWTAuth
)

func InitJWTAuth(config *cfg.Config) {
	jwtAuth = jwtauth.New(alg, []byte(config.SecretKey), nil)
}

func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtAuth
}
