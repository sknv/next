package initers

import (
	"github.com/go-chi/jwtauth"
)

const (
	alg = "HS256"
)

var (
	jwtAuth *jwtauth.JWTAuth
)

func init() {
	jwtAuth = jwtauth.New(alg, []byte(GetConfig().SecretKey), nil)
}

func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtAuth
}
