package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FirstName      string             `bson:"first_name,omitempty" json:"first_name" validate:"required"`
	LastName       string             `bson:"last_name,omitempty" json:"last_name" validate:"required"`
	Nic            string             `bson:"nic,omitempty" json:"nic" validate:"required"`
	Address        string             `bson:"address,omitempty" json:"address"`
	HomeTown       string             `bson:"home_town" json:"home_town"`
	Province       string             `bson:"province" json:"province"`
	Email          string             `bson:"email,omitempty" json:"email" validate:"required,email"`
	Nationality    string             `bson:"nationality" json:"nationality"`
	BirthDate      string             `bson:"birth_date" json:"birth_date"`
	RegisteredDate string             `bson:"registered_date" json:"registered_date"`
	PhoneNumber    string             `bson:"phone_number,omitempty" json:"phone_number" validate:"required,e164"`
}
