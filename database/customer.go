package database

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Customer ...
type Customer struct {
	ID            bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Email         string          `json:"email,omitempty" form:"email,omitempty"`
	Password      string          `json:"-" form:"-" bson:"password,omitempty"`
	FirstName     string          `json:"first_name,omitempty" form:"first_name,omitempty"`
	LastName      string          `json:"last_name,omitempty" form:"last_name,omitempty"`
	Phone         string          `json:"phone,omitempty" form:"phone,omitempty"`
	Username      string          `json:"username,omitempty" form:"username,omitempty"`
	ShippingInfos []ShippingInfo  `json:"shipping_infos,omitempty"`
	Avatar        string          `json:"avatar,omitempty"`
	Carts         []Cart          `json:"carts,omitempty"`
	Orders        []bson.ObjectId `json:"orders,omitempty" bson:"orders,omitempty"`
}

//Update ...
func (customer *Customer) Update(newCustomer *Customer) error {
	newCustomer.Email = customer.Email
	newCustomer.Password = customer.Password

	err := db.C(COL_CUSTOMERS).Update(bson.M{"_id": customer.ID}, newCustomer)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}
	return err
}

//CreateCustomer ...
func CreateCustomer(email string, password string) (Customer, error) {
	customer := Customer{
		Email:    email,
		Password: password,
	}
	var err error

	customer.ID = bson.NewObjectId()
	err = db.C(COL_CUSTOMERS).Insert(&customer)
	if err != nil {
		return Customer{}, err
	}
	return customer, err
}
