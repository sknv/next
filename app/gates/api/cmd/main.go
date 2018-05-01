package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sknv/upsale/app/gates/api/cfg"
	"github.com/sknv/upsale/app/gates/api/controllers"
	xmiddleware "github.com/sknv/upsale/app/lib/middleware"
	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	addr := cfg.GetAddr()

	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer, xmiddleware.Recoverer)

	route(router)
	xhttp.ListenAndServe(addr, router, shutdownTimeout)
}

func route(router chi.Router) {
	routeSession(router)
}

func routeSession(router chi.Router) {
	session := controllers.NewSession()
	router.Post("/login", session.Login)
}
