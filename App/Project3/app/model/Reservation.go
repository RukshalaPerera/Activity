package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reservation struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BookId              primitive.ObjectID `bson:"book_id" json:"book_id"` // Foreign key - book
	UserId              primitive.ObjectID `bson:"user_id" json:"user_id"` // Foreign key - user
	IsCompleted         bool               `bson:"is_completed" json:"is_completed"`
	EndTime             time.Time          `bson:"end_time" json:"end_time"`
	StartTime           time.Time          `bson:"start_time" json:"start_time"`
	ReservationDuration time.Duration      `bson:"reservation_duration" json:"reservation_duration"`
	Status              string             `bson:"status" json:"status"`
}

func (r *Reservation) Completed() bool {
	return r.IsCompleted
}
