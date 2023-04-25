package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,onitempty"`
	Number_account string             `json:"account_number,onitempty" bson:"account_number,onitempty,unique"`
	Amount         int64              `json:"amountDeposit,onitempty" bson:"amountDeposit,onitempty"`
	Date           time.Time          `json:"date,onitempty" bson:"date,onitempty"`
}
