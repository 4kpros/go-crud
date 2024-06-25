package data

type SignInResponse struct {
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Expires      string `json:"expires,omitempty"`
	MaxAge       string `json:"maxAge,omitempty"`
}
