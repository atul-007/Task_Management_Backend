package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Title     string             `bson:"title"`
	Priority  string             `bson:"priority"`
	DueDate   time.Time          `bson:"due_date"`
	Completed bool               `bson:"completed"`
}
