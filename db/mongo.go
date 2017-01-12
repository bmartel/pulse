package db

import (
	"strings"
	"time"
	
	"github.com/spf13/viper"
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

func Connect() *MongoStore {
	// hosts []string, username string, password string, defaultDb string
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(viper.GetString("DATABASE_HOST"), ","),
		Timeout:  60 * time.Second,
		Database: viper.GetString("DATABASE_DEFAULT"),
		Username: viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return &MongoStore{
		Name:    viper.GetString("DATABASE_DEFAULT"),
		Session: session,
	}
}
