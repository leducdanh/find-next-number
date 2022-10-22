package controller

import (
	"test-github/domain/service/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	user *user.UserService
}

func (c *Controller) NewUserController() {
	c.user = &user.UserService{}
	c.user.NewUserOrm()
}

func (c *Controller) ApiGroup(g *echo.Group) {
	c.NewUserController()
	g.GET("/users", c.user.ListAll)
}
