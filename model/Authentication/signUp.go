package Authentication

import (
	"book-store/database"

	"gopkg.in/mgo.v2/bson"
)

//SignUp ...
func SignUp(email string, password string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	customer.Email = email
	customer.Password = password
	customer.ID = bson.NewObjectId()
	database.GetMongoDB().C(database.COL_CUSTOMERS).Insert(&customer)
	return customer, err
}

//AuthReq ...
type AuthReq struct {
	Email    string `json:"email,omitempty" form:"email,omitempty"`
	Password string `json:"password,omitempty" form:"password,omiempty"`
}
