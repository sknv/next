package main

import (
	"time"

	"github.com/go-chi/chi"

	"github.com/sknv/next/app/core/initers"
	xchi "github.com/sknv/next/app/lib/chi"
	xhttp "github.com/sknv/next/app/lib/net/http"
	"github.com/sknv/next/app/rest/api/cfg"
	"github.com/sknv/next/app/rest/api/controllers"
)

const (
	concurrentRequestLimit = 1000
	requestTimeout         = 60 * time.Second
	shutdownTimeout        = 30 * time.Second
)

func main() {
	mongoSession := initers.GetMongoSession()
	defer mongoSession.Close() // Clean up.

	router := chi.NewRouter()
	xchi.UseDefaultMiddleware(router)
	xchi.ThrottleAndTimeout(router, concurrentRequestLimit, requestTimeout)
	xchi.ProvideMongoSession(router, mongoSession)

	route(router)
	xhttp.ListenAndServe(cfg.GetAddr(), router, shutdownTimeout)
}

func route(router chi.Router) {
	controllers.NewAuth().Route(router)
	controllers.NewUser().Route(router)
}
