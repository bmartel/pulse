package app

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
)

// Routes ...
type Routes struct {
	Index *Index `inject:""`
}

// Register ... Application Routes
func (r *Routes) Register(router *bone.Mux) {
	router.Get("/", http.HandlerFunc(r.Index.Get))
}

// Index ...
type Index struct {
	*render.Render `inject:""`
}

// Get ...
func (r *Index) Get(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "welcome", map[string]string{
		"title":       "Welcome to Pulse!",
		"description": "A bundle of go libs that make building impactful web apps a snap.",
	})
}
