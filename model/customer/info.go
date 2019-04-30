package customer

import (
	"book-store/database"

	"gopkg.in/mgo.v2/bson"
)

//GetCustomerByID
func GetCustomerByID(customerID string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	err = database.GetMongoDB().C(database.COL_CUSTOMERS).FindId(bson.M{"_id": bson.ObjectIdHex(customerID)}).One(&customer)
	if err != nil {
		return customer, err
	}
	return customer, err
}
