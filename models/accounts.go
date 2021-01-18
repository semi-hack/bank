package models

import (
	"time"
)

// Account ...
type Account struct {
	ID int64 `json:"id"`
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAccount struct {
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
}