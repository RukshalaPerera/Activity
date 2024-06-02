package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID       primitive.ObjectID `bson:"ID,omitempty" json:"id"`
	Filename string             `bson:"Filename,omitempty" json:"filename"`
	Title    string             `bson:"Title,omitempty" json:"title"`
}
