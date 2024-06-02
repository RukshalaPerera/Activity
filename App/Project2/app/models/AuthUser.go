package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username" json:"username" validate:"required"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Email    string             `bson:"email" json:"email" validate:"required,email"`
	RoleID   primitive.ObjectID `bson:"role_id" json:"role_id" validate:"required"`
	Role     Role               `bson:"role,omitempty" json:"role"`
}
