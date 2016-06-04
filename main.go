package main

import (
	pulse "github.com/bmartel/pulse-core/app"
	"github.com/bmartel/pulse-core/db"
	"github.com/bmartel/pulse-core/env"
	"github.com/bmartel/pulse-core/template"
)

// App ... Pulse Application Graph
type App struct {
	Routes *Routes `inject:""`
}

func main() {

	var app App

	pulse := pulse.New()

	// Database connection
	dbConn := db.New(env.DbType)

	// Content Renderer
	contentRender := template.DefaultRenderer()

	// Dependency injection
	pulse.Provide(dbConn, contentRender, &app)

	// Register routes
	pulse.Routes(app.Routes)
}
