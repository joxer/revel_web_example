package controllers

import (
	"example_app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) CreateUser(name string, username string, password string) revel.Result {
	user := models.User{Name: name, Username: username, Password: password}
	user.Insert()
	return c.RenderJSON(user)
}

func (c App) ListUser(hello string) revel.Result {
	users := models.User{}.ListUser()
	return c.RenderJSON(users)
}
