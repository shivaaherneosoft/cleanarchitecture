package repository

import "github.com/shivaaherneosoft/cleanarchitecture/api/model"

type UserRepository interface {
	GetUserDetails(string) model.UserDetails
}
