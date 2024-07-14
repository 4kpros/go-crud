package response

import "time"

type ActivateAccountResponse struct {
	ActivatedAt time.Time `json:"activatedAt"`
}
