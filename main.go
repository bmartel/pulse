package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// App ... Application Graph
type App struct {
	*app.RouteMap `inject:""`
}

func main() {

	var appGraph App

	// Database connection
	dbConn, err := gorm.Open(config.DbType, config.DbURL)
	if err != nil {
		panic(err)
	}

	// Redis Connection
	redisConn := cache.NewRedisCache(config.RedisHost, config.RedisPassword, 5*time.Minute)

	// Dependency injection
	inject.Populate(dbConn, redisConn, &appGraph)

	// Router
	r := gin.New()

	// Html Template Renderer
	r.HTMLRender = NewAmberRenderer(config.ViewDir, config.ViewExt, nil)

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
	sessionStore.Options(func() sessions.Options {
		options := sessions.Options{
			Path:     "/",
			Domain:   config.AppDomain,
			MaxAge:   86400 * 14,
			Secure:   true,
			HttpOnly: true,
		}

		if config.AppEnv != "production" {
			options.Secure = false
		}

		return options
	}())

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
