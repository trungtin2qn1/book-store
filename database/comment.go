package database

import "gopkg.in/mgo.v2/bson"

//Comment ...
type Comment struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty" form:"id,omitempty"`
	ObjectID    bson.ObjectId `json:"object_id" bson:"_object_id,omitempty"`
	ObjectType  string        `json:"object_type"`
	Content     string        `json:"content"`
	CommentDate uint64        `json:"comment_date"`
	StatusReply bool          `json:"status_reply"`
}

// //CreateCart ...
// func CreateCart(bookID bson.ObjectId) (Cart, error) {
// 	cart := Cart{}
// 	var err error
// 	return cart, err
// }

//Update ...
func (comment *Comment) Update(newCommnent Comment) error {
	return nil
}

func (comment *Comment) Delete() error {
	return nil
}
