package service

import (
	"fmt"

	"github.com/shivaaherneosoft/cleanarchitecture/api/repository"
)

type UserService interface {
	DisplayUserDetails(string)
}

type UserServiceIMPL struct {
	UserRepo repository.UserRepository
}

func GetterUserService(repo repository.UserRepository) UserServiceIMPL {
	return UserServiceIMPL{UserRepo: repo}
}

func (user *UserServiceIMPL) DisplayUserDetails(userid string) {
	userdetails := user.UserRepo.GetUserDetails(userid)
	fmt.Println("userdetails::", userdetails)
}
