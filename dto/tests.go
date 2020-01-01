package dto

import "time"

type Test struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	content   string    `json:"content"`
	tags      string    `json:"tags"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}
