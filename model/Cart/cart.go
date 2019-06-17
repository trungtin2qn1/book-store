package Cart

import (
	"book-store/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//CartsResponse ...
type CartsResponse struct {
	ID            bson.ObjectId `json:"id"`
	Amount        int           `json:"amount"`
	BookID        bson.ObjectId `json:"book_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Price         uint64        `json:"price"`
	DiscountPrice uint64        `json:"discount_price"`
}

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

	customer.Carts = append(customer.Carts, res)
	customer.Update(&customer)

	return res, err
}

//GetAllCartsInfo ...
func GetAllCartsInfo(customerID string) ([]CartsResponse, error) {
	res := []CartsResponse{}
	customer := database.Customer{
		ID: bson.ObjectIdHex(customerID),
	}
	err := database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"_id": customer.ID}).One(&customer)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(customer.Carts); i++ {
		temp := CartsResponse{}
		temp.ID = customer.Carts[i].ID
		temp.Amount = customer.Carts[i].Amount
		book := database.Book{
			ID: customer.Carts[i].BookID,
		}
		fmt.Println("book:", book)
		err = database.GetMongoDB().C(database.COL_BOOKS).Find(bson.M{"_id": book.ID}).One(&book)
		fmt.Println("book:", book)
		temp.Description = book.Description
		temp.DiscountPrice = book.DiscountPrice
		temp.Price = book.Price
		temp.Title = book.Title
		temp.BookID = book.ID
		res = append(res, temp)
	}
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
