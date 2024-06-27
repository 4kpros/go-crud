package response

type SignInResponse struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	AccessExpires  string `json:"accessExpires"`
	RefreshExpires string `json:"refreshExpires"`
	MaxAge         int    `json:"maxAge"`
}

type SignInResponse2fa struct {
	Token string `json:"accessToken"`
}
