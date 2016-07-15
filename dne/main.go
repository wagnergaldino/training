package main

import (
	"encoding/gob"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/satori/go.uuid"

	"github.com/wagnergaldino/training/dne/database"
	"github.com/wagnergaldino/training/dne/geocode"
	"github.com/wagnergaldino/training/dne/models"
	"github.com/wagnergaldino/training/dne/orm"
	"github.com/wagnergaldino/training/dne/riskzone"
	"github.com/wagnergaldino/training/dne/routers"
)

func main() {

	db := database.NewDB()
	defer db.Close()

	go riskzone.Execute(db)

	orm.Init()
	defer orm.Close()

	go geocode.Parallel(orm.Get())

	gob.Register(models.User{})
	gob.Register(models.Address{})

	middleware := negroni.Classic()
	middleware.Use(negroni.NewStatic(http.Dir("templates")))
	middleware.Use(sessions.Sessions("dne", cookiestore.New([]byte(uuid.NewV4().String()))))
	middleware.UseFunc(Authenticate)
	middleware.UseHandler(routers.GetControllers())
	middleware.Run(":8081")

}

func Authenticate(response http.ResponseWriter, request *http.Request, handler http.HandlerFunc) {
	user := sessions.GetSession(request).Get("user")
	if request.RequestURI == "/login" || user != nil {
		handler(response, request)
	} else {
		http.Redirect(response, request, "/login", http.StatusFound)
	}
}
