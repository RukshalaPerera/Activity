package model

type User struct {
	ID   int64  `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty;" json:"name"`
	Nic  string `bson:"nic,omitempty" json:"nic"`
}
