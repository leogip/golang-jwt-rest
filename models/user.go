package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Fullname  string `bson:"fullname" json:"fullname"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password,omitempty" json:"password,omitempty"`
}