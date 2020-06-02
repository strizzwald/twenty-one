package models

import "github.com/google/uuid"

type Dealer struct{
	Id uuid.UUID `json:"id"`
	Player       `json:"player"`
}
