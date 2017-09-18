package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type MGOProxy struct {
	session *mgo.Session
	db      *mgo.Database
	url     string
	dbName  string
}

var (
	Instance *MGOProxy
)

func (mgo_proxy *MGOProxy) getSession() *mgo.Session {

	session, err := mgo.Dial(mgo_proxy.url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	mgo_proxy.session = session
	mgo_proxy.db = session.DB(mgo_proxy.dbName)
	return session
}

func (mgo_proxy *MGOProxy) GetCollection(collection string) *mgo.Collection {
	return mgo_proxy.session.DB(mgo_proxy.dbName).C(collection)
}

func InitDB() {
	Instance = &MGOProxy{url: "127.0.0.1", dbName: "example"}
	Instance.getSession()
	if Instance.session.Ping() == nil {
		fmt.Println("DB INIT WELL")
	}

}
