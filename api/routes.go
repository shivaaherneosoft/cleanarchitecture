package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/cleanarchitecture/api/handler"
	"github.com/shivaaherneosoft/cleanarchitecture/api/repository"
	"github.com/shivaaherneosoft/cleanarchitecture/api/service"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// sqlrepository := repository.GetterUserSQLRepo("root", "root", "http://localhost:3306")
	mongorepository := repository.GetterUserMongoRepo("root", "root", "http://localhost:3306", "Users")
	userservice := service.GetterUserService(&mongorepository)
	userhandler := handler.GetterUserHandler(&userservice)

	router.GET("/userhandler", userhandler.DisplayUserDetails)
	return router
}
