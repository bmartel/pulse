package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

// Base ... Controller with some shared dependencies
type Base struct {
	Render *render.Render `inject:""`
	DB     *gorm.DB       `inject:""`
}
