package Authentication

import (
	"book-store/database"
	"book-store/utils"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func checkAuthData(email string, password string) bool {
	if email == "" {
		return false
	}
	if utils.ValidateFormat(email) != nil {
		return false
	}
	if len(password) < 6 {
		return false
	}
	return true
}

//SignUp ...
func SignUp(email string, password string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return customer, err
	}
	customer.Email = email
	database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"email": email}).One(&customer)
	fmt.Println("customer:", customer)
	if customer.ID.Hex() != "" {
		err = fmt.Errorf("%s", "You has registered")
		return customer, err
	}
	customer.Password, err = utils.Generate(password)
	if err != nil {
		err = fmt.Errorf("%s", "Can't hash password")
		return customer, err
	}

	customer, err = database.CreateCustomer(customer.Email, customer.Password)

	return customer, err
}

//AuthReq ...
type AuthReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
