package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/cleanarchitecture/api/service"
)

type UserHandler struct {
	UserService service.UserService
}

func GetterUserHandler(service service.UserService) UserHandler {
	return UserHandler{UserService: service}
}

func (user *UserHandler) DisplayUserDetails(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Inside UserHandler")
	user.UserService.DisplayUserDetails("USER001")
}
