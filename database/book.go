package database

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Books ...
type Book struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty" form:"id,omitempty"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Organization  string        `json:"organization"`
	Author        string        `json:"author"`
	Inventory     uint64        `json:"inventory"`
	FromTime      uint64        `json:"from_time"`
	ToTime        uint64        `json:"to_time"`
	Price         uint64        `json:"price"`
	DiscountPrice uint64        `json:"discount_price"`
	CategoryName  string        `json:"category_name"`
	Images        []string      `json:"images"`
	Comments      []Comment     `json:"comments"`
}

//CreateBook ...
func CreateBook(
	title, description, organization, author string,
	inventory, fromTime, toTime, price, discountPrice uint64,
	categoryName string, images []string, comments []Comment) (Book, error) {
	book := Book{
		Title:         title,
		Description:   description,
		Organization:  organization,
		Inventory:     inventory,
		FromTime:      fromTime,
		ToTime:        toTime,
		Price:         price,
		DiscountPrice: discountPrice,
		Author:        author,
		Images:        images,
		Comments:      comments,
	}
	var err error

	book.ID = bson.NewObjectId()
	err = db.C(COL_BOOKS).Insert(&book)
	if err != nil {
		return Book{}, err
	}
	return book, err
}

//Update ...
func (book *Book) Update(newBook Book) error {
	err := db.C(COL_BOOKS).Update(bson.M{"_id": book.ID}, newBook)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}
	return err
}

func (book *Book) Delete() error {
	return nil
}
