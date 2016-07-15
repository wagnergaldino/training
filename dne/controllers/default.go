package controllers

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"gopkg.in/unrolled/render.v1"

	"github.com/wagnergaldino/training/dne/database"
	"github.com/wagnergaldino/training/dne/orm"
	"github.com/wagnergaldino/training/dne/tools"
)

type MainController struct {
	DB     *sql.DB
	ORM    *gorm.DB
	Render *render.Render
}

func (c *MainController) Get() *MainController {
	c.DB = database.Get()
	c.ORM = orm.Get()
	c.Render = tools.GetRender()
	return c
}
