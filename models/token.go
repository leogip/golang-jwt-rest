package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Token struct
type Token struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Token  	  string `bson:"token,omitempty" json:"token,omitempty"`
	ExpriedAt time.Time `bson:"expriedAt,omitempty" json:"expriedAt,omitempty"`
}