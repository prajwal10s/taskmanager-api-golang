package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// creation of user schema

type User struct {
	ID		primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Name		string	`bson:"name" json:"name"`
	Email		string	`bson:"email" json:"email"`
	Password		string	`bson:"password" json:"password"`
	Age		int	`bson:"age,omitempty" json:"age,omitempty"`
	Tokens   []string           `bson:"tokens,omitempty" json:"tokens,omitempty"`
}