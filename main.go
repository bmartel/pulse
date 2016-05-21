package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/go-zoo/bone"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	logroni "github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/render"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// App ... Pulse Application Graph
type App struct {
	Routes *app.Routes `inject:""`
}

func main() {

	var app App

	// Database connection
	db := lib.Sqlite()
	// db := lib.Postgres()
	// db := lib.Mysql()

	// Run any necessary migrations on Application boot
	config.AutoMigrate(db)

	// Register Template Functions
	templateHelpers := lib.DefineTemplateFuncs(func(tmpl *lib.Template) {
		// Add any template functions needed
		// tmpl.Add(name, fn)
	})

	// Content Renderer
	rend := render.New(render.Options{
		IndentJSON: true,
		Layout:     "layout",
		Extensions: []string{".html"},
		Funcs:      templateHelpers,
	})

	// Dependency Injection of shared services
	inject.Populate(db, rend, &app)

	// Initialize Router
	router := bone.New()

	// Initialize Middleware
	mw := negroni.New()

	// Compression
	mw.Use(gzip.Gzip(gzip.DefaultCompression))

	// Public Files ... Assets, Images
	mw.Use(negroni.NewStatic(http.Dir(lib.AssetDir)))

	// Recovery
	mw.Use(negroni.NewRecovery())

	// Logging
	mw.Use(logroni.NewMiddleware())

	// Session Store
	store := cookiestore.New([]byte(lib.AppKey))
	mw.Use(sessions.Sessions(lib.SessionKey, store))

	// Csrf
	csrfProtect := csrf.Protect(
		[]byte(lib.AppKey),
		csrf.Secure(lib.AppEnv == "production"),
	)

	// Register Application Routes
	app.Routes.Register(router)

	// Apply Middleware to Router
	mw.UseHandler(context.ClearHandler(csrfProtect(router)))

	// Run Application
	mw.Run(lib.AppPort)
}
