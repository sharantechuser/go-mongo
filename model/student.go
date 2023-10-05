package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Address  Address            `json:"address" bson:address`
}

type Address struct {
	Address1 string `json:"address" bson:"address"`
	City     string `json:"city" bson:"city"`
	Pin      string `json:"pin" bson:"pin"`
}
