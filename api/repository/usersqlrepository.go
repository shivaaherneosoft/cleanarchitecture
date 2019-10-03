package repository

import (
	"fmt"

	"github.com/shivaaherneosoft/cleanarchitecture/api/model"
)

type UserSQLRepo struct {
	DBUser string
	DBPass string
	DBUrl  string
}

func GetterUserSQLRepo(dbuser, dbpass, dburl string) UserSQLRepo {
	return UserSQLRepo{DBUser: dbuser, DBPass: dbpass, DBUrl: dburl}
}

func (user *UserSQLRepo) GetUserDetails(userid string) model.UserDetails {
	fmt.Println("SQL Connection::", user)
	userdetails := model.UserDetails{UserId: userid, UserName: "Shiva"}
	return userdetails
}
