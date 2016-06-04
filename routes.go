package main

import (
	"github.com/bmartel/pulse-core/controller"
	"github.com/bmartel/pulse-core/router"
)

// Routes ...
type Routes struct {
	*controller.Index `inject:""`
}

// Register ... Application Routes
func (r *Routes) Register(router router.Router) {
	router.GetFunc("/", r.Index.Get)
}
