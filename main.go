package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// App ... Application Graph
type App struct {
	*app.RouteMap `inject:""`
}

func main() {

	var appGraph App

	// Database connection
	mongoConn := db.ConnectMongo(strings.Split(config.DbHost, ","), config.DbUser, config.DbPass, config.DbDefault)

	// Redis Connection
	redisConn := cache.NewRedisCache(config.RedisHost, config.RedisPassword, 5*time.Minute)

	// Dependency injection
	inject.Populate(mongoConn, redisConn, &appGraph)

	// Router
	r := gin.New()

	// Html Template Renderer
	r.HTMLRender = NewAmberRenderer(config.ViewDir, config.ViewExt, nil)

	// Recover from system panics
	r.Use(gin.Recovery())

	// Compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Load Assets
	r.Static(config.AssetPath, config.AssetDir)

	// Declare a subgroup for app routes so middleware is not run for templates or assets
	router := r.Group("/")

	// Logging
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	// Cookie Session Store
	sessionStore := sessions.NewCookieStore([]byte(config.AppKey))
	sessionStore.Options(sessions.Options{
		Path:     "/",
		Domain:   config.AppDomain,
		MaxAge:   86400 * 14,
		Secure:   config.AppEnv == "production",
		HttpOnly: true,
	})

	// Session
	router.Use(sessions.Sessions(config.SessionKey, sessionStore))

	// CSRF
	router.Use(csrf.Middleware(csrf.Options{
		Secret: config.AppKey,
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// Register routes
	appGraph.RouteMap.Register(router)

	// Run application on $APP_PORT
	r.Run(config.AppPort)
}
