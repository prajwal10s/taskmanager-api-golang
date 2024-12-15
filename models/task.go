package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//creation of task schema

type Task struct {
	ID		primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Description		string	`bson:"description" json:"description"`
	Completed		bool	`bson:"completed" json:"completed"`
	Owner		primitive.ObjectID	`bson:"owner" json:"owner"`
}