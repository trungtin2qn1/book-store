package database

import "gopkg.in/mgo.v2/bson"

//Customer ...
type Customer struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Email     string        `json:"email,omitempty" form:"email,omitempty"`
	Password  string        `json:"password,omitempty" form:"password,omitempty"`
	FirstName string        `json:"first_name,omitempty" form:"first_name,omitempty"`
	LastName  string        `json:"last_name,omitempty" form:"last_name,omitempty"`
}
