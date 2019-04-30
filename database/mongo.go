package database

import (
	"book-store/utils"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const (
	COL_CUSTOMERS = "customers"
)

//InitMongo ...
func InitMongo() {

	info := &mgo.DialInfo{
		Addrs:    []string{utils.GetEnv("MONGO_HOST", "localhost:27017")},
		Timeout:  60 * time.Second,
		Database: utils.GetEnv("MONGO_DATABASE", "book_store"),
		Username: utils.GetEnv("MONGO_USERNAME", "admin"),
		Password: utils.GetEnv("MONGO_PASSWORD", "bs_1234"),
	}

	fmt.Println(info)

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	session.Ping()

	db = session.DB(utils.GetEnv("MONGO_DATABASE", "book_store"))
}

//GetMongoDB ...
func GetMongoDB() *mgo.Database {
	return db
}

//Get ...
func Get() {
	customers := []Customer{}
	err := db.C("test").Find(bson.M{}).All(&customers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(customers)
}
