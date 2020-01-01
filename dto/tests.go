package dto

import "time"

type Test struct {
	Id         int64
	Name       string
	Thumbnail  string
	content    string
	tags       string
	created_at time.Time
	updated_at time.Time
}
