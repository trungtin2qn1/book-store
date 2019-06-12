package Book

import (
	"book-store/database"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//CreateBook ...
func CreateBook(book *database.Book) (database.Book, error) {
	res := database.Book{}
	res, err := database.CreateBook(
		book.Title,
		book.Description,
		book.Organization,
		book.Author,
		book.Inventory,
		book.FromTime,
		book.ToTime,
		book.Price,
		book.DiscountPrice,
		book.CategoryName,
		book.Images,
		book.Comments)
	return res, err
}

//GetAllBooks ...
func GetAllBooks() ([]database.Book, error) {
	res := []database.Book{}
	err := database.GetMongoDB().C(database.COL_BOOKS).Find(bson.D{}).All(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

//GetBookByID ...
func GetBookByID(bookID string) (database.Book, error) {
	fmt.Println("book id:", bookID)
	res := database.Book{}
	res.ID = bson.ObjectIdHex(bookID)
	err := database.GetMongoDB().C(database.COL_BOOKS).Find(bson.M{"_id": res.ID}).One(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}
