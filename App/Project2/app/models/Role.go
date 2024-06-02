package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	RoleName string             `bson:"roleName" json:"roleName"`
}

func NewRole(roleName string) *Role {
	return &Role{
		RoleName: roleName,
	}
}

func CreateDefaultRoles() []*Role {
	defaultRoles := []*Role{
		NewRole("user"),
		NewRole("moderator"),
		NewRole("admin"),
	}
	return defaultRoles
}
