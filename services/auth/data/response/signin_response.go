package response

import "time"

type SignInResponse struct {
	AccessToken string    `json:"accessToken"`
	Expires     time.Time `json:"expires"`
}

type SignInResponse2fa struct {
	Token string `json:"token"`
}
