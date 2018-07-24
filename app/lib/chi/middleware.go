package chi

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_chi"
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	lib "github.com/sknv/next/app/lib/middleware"
	mongo "github.com/sknv/next/app/lib/mongo/middleware"
)

func UseDefaultMiddleware(router chi.Router) {
	router.Use(
		middleware.RealIP, middleware.Logger, middleware.Recoverer, lib.Recoverer,
	)
}

func UseThrottle(router chi.Router, concurrentRequestLimit int) {
	router.Use(middleware.Throttle(concurrentRequestLimit))
}

func UseTimeout(router chi.Router, requestTimeout time.Duration) {
	router.Use(middleware.Timeout(requestTimeout))
}

func UseLimitHandler(router chi.Router, requestLimit float64) {
	router.Use(tollbooth_chi.LimitHandler(tollbooth.NewLimiter(requestLimit, nil)))
}

func WithLimitHandler(router chi.Router, requestLimit float64) chi.Router {
	return router.With(tollbooth_chi.LimitHandler(tollbooth.NewLimiter(requestLimit, nil)))
}

func ProvideMongoSession(router chi.Router, session *mgo.Session) {
	router.Use(mongo.WithMongoSession(session))
}
