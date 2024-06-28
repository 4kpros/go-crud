package types

import "time"

type JwtToken struct {
	UserId  uint
	Role    string
	Expires time.Time
	Device  string
}
