package types

type Filter struct {
	Search  string `json:"search"`
	OrderBy string `json:"orderBy"`
	Sort    string `json:"sort"`
}
