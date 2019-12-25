package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/phamtanlong/go-crud/news/pb"
	"log"
	"time"
)

type News struct {
	gorm.Model
	Title     string
	Thumbnail string
	Content   string
	Tags      string
	CreatedBy *uint
	UpdatedBy *uint
}
type NewsService struct {
	DB *gorm.DB
	AuthenClient pb.AuthenticationClient
}

func (n NewsService) Create(ctx context.Context, token string, title string, thumbnail string, content string, tags string) (uint, error) {
	log.Printf("> create \ntitle: %s\nthumbnail: %s\ncontent: %s\ntags: %s", title, thumbnail, content, tags)

	res, err := n.AuthenClient.Verify(ctx, &pb.VerifyRequest{Token:token})
	if err != nil {
		return 0, ErrorString{Message:"can not connect authen service"}
	}
	if len(res.Error) > 0 {
		return 0, ErrorString{Message:res.Error}
	}
	userId := uint(res.Id)

	record := News{
		Title:     title,
		Thumbnail: thumbnail,
		Content:   content,
		Tags:      tags,
		CreatedBy: &userId,
	}

	n.DB.Create(&record)
	return record.ID, nil
}
func (n NewsService) Update(ctx context.Context, token string, id uint, title string, thumbnail string, content string, tags string) error {
	log.Printf("> update id %d\ntitle: %s\nthumbnail: %s\ncontent: %s\ntags: %s", id, title, thumbnail, content, tags)

	res, err := n.AuthenClient.Verify(ctx, &pb.VerifyRequest{Token:token})
	if err != nil {
		return ErrorString{Message:"can not connect authen service"}
	}
	if len(res.Error) > 0 {
		return ErrorString{Message:res.Error}
	}
	userId := uint(res.Id)

	var record News
	if n.DB.First(&record, id).RecordNotFound() {
		return ErrorString{Message:"record not found"}
	}

	record.Title = title
	record.Thumbnail = thumbnail
	record.Content = content
	record.Tags = tags
	record.UpdatedBy = &userId

	n.DB.Save(&record)
	return nil
}
func (n NewsService) Delete(ctx context.Context, token string, id uint) error {
	log.Printf("> delete id %d", id)

	res, err := n.AuthenClient.Verify(ctx, &pb.VerifyRequest{Token:token})
	if err != nil {
		return ErrorString{Message:"can not connect authen service"}
	}
	if len(res.Error) > 0 {
		return ErrorString{Message:res.Error}
	}
	userId := uint(res.Id)

	var record News
	if n.DB.First(&record, id).RecordNotFound() {
		return ErrorString{Message:"record not found"}
	}

	record.UpdatedBy = &userId

	now := time.Now()
	record.DeletedAt = &now
	n.DB.Save(&record)

	//n.DB.Where("id = ?", id).Delete(&News{})
	return nil
}
func (n NewsService) Read(ctx context.Context, token string, id uint) (*News, error) {
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