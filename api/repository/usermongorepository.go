package repository

import (
	"fmt"

	"github.com/shivaaherneosoft/cleanarchitecture/api/model"
)

type UserMongoRepo struct {
	DBUser       string
	DBPass       string
	DBUrl        string
	DBCollection string
}

func GetterUserMongoRepo(dbuser, dbpass, dburl, dbcollection string) UserMongoRepo {
	return UserMongoRepo{DBUser: dbuser, DBPass: dbpass, DBUrl: dburl, DBCollection: dbcollection}
}

func (user *UserMongoRepo) GetUserDetails(userid string) model.UserDetails {
	fmt.Println("Mongo Connection::", user)
	userdetails := model.UserDetails{UserId: userid, UserName: "Shiva"}
	return userdetails
}
