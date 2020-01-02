package dto

import (
	"fmt"
	"time"
)

type Test struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (test Test) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\n:Thumbnail: %s\nContent: %s\nTags: %s\nCreateAt: %s\nUpdatedAt: %s\n", test.Id, test.Name, test.Thumbnail, test.Content, test.Tags, test.CreatedAt, test.UpdatedAt)
}
