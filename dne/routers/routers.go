package routers

import (
	"github.com/julienschmidt/httprouter"

	"github.com/wagnergaldino/training/dne/controllers"
)

func GetControllers() (router *httprouter.Router) {
	router = httprouter.New()

	login := &controllers.LoginController{}
	home := &controllers.HomeController{}
	address := &controllers.AddressController{}

	router.GET("/login", login.Form)
	router.POST("/login", login.Login)
	router.GET("/logout", login.Logout)
	router.GET("/", home.Home)
	router.GET("/home", home.Home)
	router.POST("/address", address.Search)

	return
}
