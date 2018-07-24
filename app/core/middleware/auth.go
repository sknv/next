package middleware

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/sknv/next/app/core/globals"
	"github.com/sknv/next/app/services"
)

func RequireLogin(router chi.Router) {
	jwtAuth := globals.GetJWTAuth()
	router.Use(
		// Require presence of valid JWT.
		jwtauth.Verifier(jwtAuth), jwtauth.Authenticator,
		// Require presence of current user.
		CurrentUserVerifier,
	)
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

func CurrentUserVerifier(next http.Handler) http.Handler {
	whoami := services.NewWhoAmI()
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, err := whoami.GetCurrentUser(r)
		if err != nil {
			log.Print("[ERROR] verify current user: ", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
