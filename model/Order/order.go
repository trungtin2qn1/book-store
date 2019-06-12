package Order

import (
	"book-store/database"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//GetAllOrders ...
func GetAllOrders(customerID string) ([]database.Order, error) {
	res := []database.Order{}
	err := database.GetMongoDB().C(database.COL_ORDERS).Find(bson.D{}).All(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

//CreateOrder ...
func CreateOrder(note string, customerID string) (database.Order, error) {
	customer := database.Customer{}
	customer.ID = bson.ObjectIdHex(customerID)
	err := database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"_id": customer.ID}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}

	bookIDs := []bson.ObjectId{}

	for i := 0; i < len(customer.Carts); i++ {
		bookIDs = append(bookIDs, customer.Carts[i].BookID)
	}

	res, _ := database.CreateOrder(note,
		uint64(time.Now().Unix()),
		10000,
		0,
		bookIDs,
		database.Payment{},
		customer.ID)

	customer.Carts = []database.Cart{}

	customer.Update(&customer)

	return res, nil
}
