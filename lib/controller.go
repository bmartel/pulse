package lib

import (
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

// BaseController a Controller with some shared dependencies
type BaseController struct {
	*render.Render `inject:""`
	*gorm.DB       `inject:""`
}
