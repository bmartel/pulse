package main // import "github.com/bmartel/pulse"

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/bmartel/gin-amber"
	"github.com/bmartel/pulse/app"
	"github.com/bmartel/pulse/config"
	"github.com/bmartel/pulse/db"
	"github.com/bmartel/zero"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/contrib/cache"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/utrack/gin-csrf"
)

// App ... Application Graph
type App struct {
	*app.RouteMap `inject:""`
}

func main() {
		// Apply environment and config file
	config.Apply()

	// Start app with db connection on specified port
	AppStart(viper.GetInt("PORT"), db.Connect())
}

// AppStart will run the application with the specified db connection and on specified port
func AppStart(port int, conn *db.MongoStore) {
	var appGraph App

	// Redis Connection
	redisConn := cache.NewRedisCache(viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PASSWORD"), 5*time.Minute)

	// Validation
	zeroValidator := zero.New("valid")

	// Dependency injection
	inject.Populate(conn, redisConn, zeroValidator, &appGraph)

	// Router
	r := gin.New()

	// Html Template Renderer
	r.HTMLRender = ginamber.NewViewRenderer(viper.GetString("VIEW_DIR"), viper.GetString("VIEW_EXT"), nil)

	// Recover from system panics
	r.Use(gin.Recovery())

	// Compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Load Assets
	r.Static(viper.GetString("ASSET_PATH"), viper.GetString("ASSET_DIR"))

	// Declare a subgroup for app routes so middleware is not run for templates or assets
	router := r.Group("/")

	// Logging
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	// Cookie Session Store
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("APP_KEY")))
	sessionStore.Options(sessions.Options{
		Path:     "/",
		Domain:   viper.GetString("APP_DOMAIN"),
		MaxAge:   86400 * 14,
		Secure:   viper.GetString("APP_ENV") == "production",
		HttpOnly: true,
	})

	// Session
	router.Use(sessions.Sessions(viper.GetString("SESSION_KEY"), sessionStore))

	// CSRF
	router.Use(csrf.Middleware(csrf.Options{
		Secret: viper.GetString("APP_KEY"),
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// Register routes
	appGraph.RouteMap.Register(router)

	// Run application on $PORT
	r.Run(":" + strconv.Itoa(port))
}
