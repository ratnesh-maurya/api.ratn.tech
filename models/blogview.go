package models

type BlogView struct {
	Slug  string `json:"slug" bson:"slug"`
	Views int    `json:"views" bson:"views"`
}
