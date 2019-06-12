package Cart

import (
	"book-store/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//AddBookToCart ...
func AddBookToCart(customerID string, cart database.Cart) (database.Cart, error) {
	res := database.Cart{}
	res, err := database.CreateCart(cart.Amount, cart.BookID)
	customer := database.Customer{}
	customer.ID = bson.ObjectIdHex(customerID)
	err = database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}

	customer.Carts = append(customer.Carts, cart)
	customer.Update(&customer)

	return res, err
}

//GetAllCartsInfo ...
func GetAllCartsInfo(customerID string) ([]database.Cart, error) {
	res := []database.Cart{}
	customer := database.Customer{}
	err := database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.D{}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}
	res = customer.Carts
	return res, err
}

//DeleteProductFromCart ...
func DeleteProductFromCart(customerID, cartID string) error {
	customer := database.Customer{}
	customer.ID = bson.ObjectIdHex(customerID)
	err := database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(customer.Carts); i++ {
		if customer.Carts[i].ID.Hex() == cartID {
			customer.RemoveCart(i)
			break
		}
	}

	customer.Update(&customer)
	return err
}

//UpdateCartInfo ...
func UpdateCartInfo(customerID, cartID string, newCart database.Cart) (database.Cart, error) {
	res := database.Cart{}

	customer := database.Customer{}
	customer.ID = bson.ObjectIdHex(customerID)
	err := database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"_id": customer.ID}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(customer.Carts); i++ {
		if customer.Carts[i].ID.Hex() == cartID {
			customer.Carts[i].Amount = newCart.Amount
			res = customer.Carts[i]
			break
		}
	}

	customer.Update(&customer)
	return res, err
}
