package controllers

import "net/http"

// Welcome ... Controller to render a sample welcome page
type Welcome struct {
	*Base `inject:""`
}

// Index ... The Application index page
func (c *Welcome) Index(w http.ResponseWriter, req *http.Request) {
	c.Render.HTML(w, http.StatusOK, "welcome", map[string]string{
		"title":       "Welcome to Pulse!",
		"description": "A bundle of go libs that make building impactful web apps a snap.",
	})
}
