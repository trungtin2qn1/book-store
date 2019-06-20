package database

import "gopkg.in/mgo.v2/bson"

//Staff ...
type Staff struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty" form:"id,omitempty"`
	Email     string        `json:"email" form:"email,omitempty"`
	Password  string        `json:"-" form:"-" bson:"password,omitempty"`
	FirstName string        `json:"first_name" form:"first_name,omitempty"`
	LastName  string        `json:"last_name" form:"last_name,omitempty"`
	Phone     string        `json:"phone" form:"phone,omitempty"`
	Username  string        `json:"username" form:"username,omitempty"`
	Position  string        `json:"position"`
	Avatar    string        `json:"avatar"`
}

// //CreateStaff ...
// func CreateStaff(bookID bson.ObjectId) (Cart, error) {
// 	cart := Cart{}
// 	var err error
// 	return cart, err
// }

//Update ...
func (staff *Staff) Update(newStaff Staff) error {
	return nil
}

//Delete
func (staff *Staff) Delete() error {
	return nil
}
