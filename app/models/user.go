package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name,omitempty" json:"name"`
	Email           string             `bson:"email,omitempty" json:"email"`
	EmailVerifiedAt *time.Time         `bson:"email_verified_at" json:"email_verified_at"`
	Status          string             `bson:"status" json:"status"`
	Password        string             `bson:"password,omitempty" json:"-"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
