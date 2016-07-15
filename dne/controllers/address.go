package controllers

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/wagnergaldino/training/dne/models"
)

type AddressController struct {
	MainController
}

type AddressMessage struct {
	Message string
}

type AddressResponse struct {
	models.User
	models.Address
	AddressMessage
}

func (c *AddressController) Search(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	message := AddressMessage{"Search result for Postal Code"}

	user := sessions.GetSession(request).Get("user").(models.User)

	address := models.Address{
		Cep: request.FormValue("cep"),
	}

	err := c.Get().ORM.Where(&address).Find(&address).Error

	if err == nil {
		responseType := request.FormValue("response_type")

		if responseType == "xml" {
			c.Get().Render.HTML(response, http.StatusOK, "address_xml", AddressResponse{user, address, message})
		} else if responseType == "json" {
			c.Get().Render.HTML(response, http.StatusOK, "address_json", AddressResponse{user, address, message})
		} else {
			c.Get().Render.HTML(response, http.StatusOK, "address_html", AddressResponse{user, address, message})
		}

		return
	}

	if err == gorm.ErrRecordNotFound {
		message.Message = "Search result for Postal Code " + request.FormValue("cep") + ": ADDRESS NOT FOUND"
		c.Get().Render.HTML(response, http.StatusFound, "home", AddressResponse{user, models.Address{}, message})
		return
	}

	if err != nil {
		c.Get().Render.HTML(response, http.StatusInternalServerError, "500", nil)
	}

}
