package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	Name    string
	Session *mgo.Session
}

func (conn *MongoStore) Close() {
	conn.Session.Close()
}

func (conn *MongoStore) Open() *mgo.Session {
	return conn.Session.Copy()
}

func ConnectMongo(hosts []string, username string, password string, defaultDb string) *MongoStore {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    hosts,
		Timeout:  60 * time.Second,
		Database: defaultDb,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return &MongoStore{
		Name:    defaultDb,
		Session: session,
	}
}
