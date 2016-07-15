package controllers

import (
	"net/http"
	"strings"

	"github.com/goincremental/negroni-sessions"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/wagnergaldino/training/dne/models"
)

type LoginController struct {
	MainController
}

type LoginMessage struct {
	Message string
}

func (c *LoginController) Form(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	message := LoginMessage{"Sign In"}

	if strings.Contains(request.Referer(), "/login") {
		message.Message = "Sign In - Invalid Username/Password"
	}

	c.Get().Render.HTML(response, http.StatusOK, "login", message)
}

func (c *LoginController) Login(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	user := &models.User{
		Usuario: request.FormValue("usuario"),
		Senha:   request.FormValue("senha"),
	}

	err := c.Get().ORM.Where(user).Find(user).Error

	if err == nil {
		session := sessions.GetSession(request)
		session.Set("user", user)

		http.Redirect(response, request, "/home", http.StatusFound)
		return
	}

	if err == gorm.ErrRecordNotFound {
		http.Redirect(response, request, "/login", http.StatusFound)
		return
	}

	if err != nil {
		c.Get().Render.HTML(response, http.StatusInternalServerError, "500", nil)
	}

}

func (c *LoginController) Logout(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	session := sessions.GetSession(request)
	session.Delete("user")

	http.Redirect(response, request, "/login", http.StatusFound)
	return
}
