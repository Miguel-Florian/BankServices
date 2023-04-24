package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,onitempty"`
	Number_account string             `json:"account_number,onitempty" bson:"account_number,onitempty,unique"`
	Type_account   string             `json:"account_type,onitempty" bson:"account_type,onitempty"`
	Amount         int64              `json:"amount,onitempty" bson:"amount,onitempty"`
	DateCreation   time.Time          `json:"dateCreation,onitempty" bson:"dateCreation,onitempty"`
	DateUpdate     time.Time          `json:"dateUpdate,onitempty" bson:"dateUpdate,onitempty"`
	FirstName      string             `json:"firstname,onitempty" bson:"firstname,onitempty"`
	LastName       string             `json:"lastname,onitempty,unique" bson:"lastname,onitempty"`
	Email          string             `json:"email,onitempty" bson:"email,onitempty,unique"`
	Adresse        string             `json:"adresse,onitempty" bson:"adresse,onitempty,unique"`
	Phone          string             `json:"phone,omnitempty" bson:"phone,onitempty,unique"`
	Password       string             `json:"-" bson:"password,onitempty"`
}
