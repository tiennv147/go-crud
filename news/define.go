package news

import (
	"context"
	"github.com/jinzhu/gorm"
)

// NewsService is a service to serve news
type NewsService interface {
	Create(ctx context.Context, title string, thumbnail string, content string, tags string) (id uint, err error)
	Update(ctx context.Context, id uint, title string, thumbnail string, content string, tags string) error
	Delete(ctx context.Context, id uint) error
	Read(ctx context.Context, id uint) (*News, error)
}

// News is base model for this module
type News struct {
	gorm.Model
	Title     string
	Thumbnail string
	Content   string
	Tags      string
}

