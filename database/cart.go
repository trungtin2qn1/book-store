package database

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Cart ...
type Cart struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty" form:"id,omitempty"`
	Amount int           `json:"amount"`
	BookID bson.ObjectId `json:"book_id" bson:"_book_id,omitempty" form:"book_id,omitempty"`
}

//CreateCart ...
func CreateCart(amount int, bookID bson.ObjectId) (Cart, error) {
	cart := Cart{
		Amount: amount,
		BookID: bookID,
	}
	cart.ID = bson.NewObjectId()
	fmt.Println(cart)
	return cart, nil
}
