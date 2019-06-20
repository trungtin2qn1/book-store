package Customer

import (
	"book-store/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//GetCustomerByID
func GetCustomerByID(customerID string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	customer.ID = bson.ObjectIdHex(customerID)
	err = database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"_id": customer.ID}).One(&customer)
	if err != nil {
		fmt.Println(err)
		return customer, err
	}
	return customer, err
}

//UpdateCustomerInfo ...
func UpdateCustomerInfo(customer *database.Customer, newCustomer *database.Customer) error {
	return customer.Update(newCustomer)
}
