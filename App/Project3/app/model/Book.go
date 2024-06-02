package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title           string             `bson:"title,omitempty" json:"title"`
	Author          string             `bson:"author,omitempty" json:"author"`
	BookName        string             `bson:"book_name,omitempty" json:"book_name"`
	IsBookAvailable bool               `bson:"is_book_available,omitempty" json:"is_book_available"`
	RegisterDate    string             `bson:"register_date,omitempty" json:"register_date"`
}
