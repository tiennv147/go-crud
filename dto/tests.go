package dto

import (
	_ "time"

	"github.com/jinzhu/gorm"
)

type Test struct {
	gorm.Model
	// Id        int64     `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Content   string `json:"content"`
	Tags      string `json:"tags"`
	// createdAt time.Time `json:"created_at"`
	// updatedAt time.Time `json:"updated_at"`
}
