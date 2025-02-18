package models

import "time"

type Post struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content string `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Written_by int `json:"written_by"`
}