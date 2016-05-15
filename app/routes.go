package app

import (
	"net/http"

	"github.com/bmartel/pulse/app/controllers"
	"github.com/go-zoo/bone"
)

// Routes ...
type Routes struct {
	Welcome *controllers.Welcome `inject:""`
}

// Register ... Application Routes
func (r *Routes) Register(router *bone.Mux) {
	router.Get("/", http.HandlerFunc(r.Welcome.Index))
}
