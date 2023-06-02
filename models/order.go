package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID  `bson:"_id"`
	Category  *string             `json:"category"`
    Dish      *string             `json:"dish"`
	Quantity  *int                `json:"quantity"`
	TableNumber  *int             `json:"tablenumber"`
	ServerName   *string          `json:"servername"` 
	Price     *float64            `json:"price"`
	  
}