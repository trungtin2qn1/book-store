package database

import "gopkg.in/mgo.v2/bson"

//Order ...
type Order struct {
	ID          bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Note        string          `json:"note,omitempty"`
	OrderDate   uint64          `json:"order_date,omitempty"`
	ShippingFee uint64          `json:"shipping_fee,omitempty"`
	Status      uint64          `json:"status,omitempty"`
	BookIDs     []bson.ObjectId `json:"book_id,omitempty" bson:"_book_id,omitempty"`
	Payment     Payment         `json:"payment,omitempty"`
	CustomerID  bson.ObjectId   `json:"customer_id,omitempty" bson:"_customer_id,omitempty"`
}

func CreateOrder(
	note string,
	orderdate, shippingFee, status uint64,
	bookIDs []bson.ObjectId, payment Payment,
	customerID bson.ObjectId) (Order, error) {
	order := Order{
		Note:        note,
		OrderDate:   orderdate,
		ShippingFee: shippingFee,
		Status:      status,
		BookIDs:     bookIDs,
		Payment:     payment,
		CustomerID:  customerID,
		ID:          bson.NewObjectId(),
	}
	err := db.C(COL_ORDERS).Insert(&order)
	if err != nil {
		return Order{}, err
	}
	return order, err
}

//GetAllOrders...
func GetAllOrders() ([]Order, error) {
	orders := []Order{}

	return orders, nil
}
