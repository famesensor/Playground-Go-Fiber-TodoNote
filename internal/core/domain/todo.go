package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Status      string             `json:"status" bson:"status" validate:"required"`
	Priority    int                `json:"priority" bson:"priority" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}
