package data

import "time"

type ValidateAccountResponse struct {
	ValidatedAt time.Time `json:"validatedAt"`
}
