package controllers

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"

	"github.com/wagnergaldino/training/dne/models"
)

type HomeController struct {
	MainController
}

type HomeMessage struct {
	Message string
}

type HomeResponse struct {
	models.User
	HomeMessage
}

func (c *HomeController) Home(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	message := HomeMessage{"Search"}

	session := sessions.GetSession(request)
	user := session.Get("user").(models.User)

	c.Get().Render.HTML(response, http.StatusOK, "home", HomeResponse{user, message})
}
