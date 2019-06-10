package database

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//Books ...
type Book struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	Organization  string        `json:"organization,omitempty"`
	Author        string        `json:"author,omitempty"`
	Inventory     uint64        `json:"inventory,omitempty"`
	FromTime      uint64        `json:"from_time,omitempty"`
	ToTime        uint64        `json:"to_time,omitempty"`
	Price         uint64        `json:"price,omitempty"`
	DiscountPrice uint64        `json:"discount_price,omitempty"`
	CategoryName  string        `json:"category_name,omitmpey"`
	Images        []string      `json:"images,omitempty"`
	Comments      []Comment     `json:"comments,omitempty"`
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
