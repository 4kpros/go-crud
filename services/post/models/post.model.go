package models

import "github.com/4kpros/go-crud/common/types"

type Post struct {
	types.BaseGormModel
	Title string `json:"title"`
	Body  string `json:"body"`
}
