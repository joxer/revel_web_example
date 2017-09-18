package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

func (u *User) Insert() bool {
	collection := Instance.GetCollection("User")
	err := collection.Insert(u)

	if err != nil {
		panic(err)
	}
	return true
}

func (u User) ListUser() []User {


	var users []User
	collection := Instance.GetCollection("User")
	defer collection.Database.Session.Close()
	collection.Find(bson.M{}).Select(bson.M{}).All(&users)
	return users
}
