package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
)

type News struct {
	gorm.Model
	Title     string
	Thumbnail string
	Content   string
	Tags      string
}
type NewsService struct {
	DB *gorm.DB
}

func (n NewsService) Create(ctx context.Context, title string, thumbnail string, content string, tags string) (uint, error) {
	log.Printf("> create \ntitle: %s\nthumbnail: %s\ncontent: %s\ntags: %s", title, thumbnail, content, tags)

	record := News{
		Title:     title,
		Thumbnail: thumbnail,
		Content:   content,
		Tags:      tags,
	}

	n.DB.Create(&record)
	return record.ID, nil
}
func (n NewsService) Update(ctx context.Context, id uint, title string, thumbnail string, content string, tags string) error {
	log.Printf("> update id %d\ntitle: %s\nthumbnail: %s\ncontent: %s\ntags: %s", id, title, thumbnail, content, tags)

	var record News
	if n.DB.First(&record, id).RecordNotFound() {
		return ErrorString{Message:"record not found"}
	}

	record.Title = title
	record.Thumbnail = thumbnail
	record.Content = content
	record.Tags = tags

	n.DB.Save(&record)
	return nil
}
func (n NewsService) Delete(ctx context.Context, id uint) error {
	log.Printf("> delete id %d", id)

	n.DB.Where("id = ?", id).Delete(&News{})
	return nil
}
func (n NewsService) Read(ctx context.Context, id uint) (*News, error) {
	log.Printf("> read id %d", id)
	
	var record News
	if n.DB.First(&record, id).RecordNotFound() {
		return nil, ErrorString{Message:"record not found"}
	}
	return &record, nil
}

// errorString is a trivial implementation of error.
type ErrorString struct {
	Message string
}

func (e ErrorString) Error() string {
	return e.Message
}