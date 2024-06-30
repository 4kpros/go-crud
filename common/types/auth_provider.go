package types

type AuthProvider struct {
	Key         string
	Name        string
	Description string
}

var AuthProviders = map[string]bool{
	"google":   true,
	"facebook": true,
}
