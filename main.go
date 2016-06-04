package main

import (
	pulseApp "github.com/bmartel/pulse-core/app"
	"github.com/bmartel/pulse-core/db"
	"github.com/bmartel/pulse-core/env"
	"github.com/bmartel/pulse-core/template"
	"github.com/facebookgo/inject"
)

// App ... Pulse Application Graph
type App struct {
	*Routes `inject:""`
}

func main() {

	var appGraph App

	pulse := pulseApp.New()

	// Database connection
	dbConn := db.New(env.DbType)

	// Content Renderer
	contentRender := template.DefaultRenderer()

	// Dependency injection
	inject.Populate(dbConn, contentRender, &appGraph)

	// Register routes
	pulse.Routes(appGraph.Routes)

	// Run application on $APP_PORT
	pulse.Run()
}
