package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/phamtanlong/go-crud/news/service"
)

type CreateRequest struct {
	Token     string
	Title     string
	Thumbnail string
	Content   string
	Tags      string
}
type CreateResponse struct {
	Id    uint
	Error error
}

func MakeCreateEndpoint(n service.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := n.Create(ctx, req.Token, req.Title, req.Thumbnail, req.Content, req.Tags)
		return CreateResponse{Id: id, Error: err}, nil
	}
}

type UpdateRequest struct {
	Token     string
	Id        uint
	Title     string
	Thumbnail string
	Content   string
	Tags      string
}
type UpdateResponse struct {
	Error error
}

func MakeUpdateEndpoint(n service.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := n.Update(ctx, req.Token, req.Id, req.Title, req.Thumbnail, req.Content, req.Tags)
		return UpdateResponse{Error: err}, nil
	}
}

type DeleteRequest struct {
	Token string
	Id    uint
}
type DeleteResponse struct {
	Error error
}

func MakeDeleteEndpoint(n service.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := n.Delete(ctx, req.Token, req.Id)
		return DeleteResponse{Error: err}, nil
	}
}

type ReadRequest struct {
	Token string
	Id    uint
}
type ReadResponse struct {
	News  *service.News
	Error error
}

func MakeReadEndpoint(n service.NewsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadRequest)
		N, err := n.Read(ctx, req.Token, req.Id)
		return ReadResponse{News: N, Error: err}, nil
	}
}

type Endpoints struct {
	Create endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
	Read   endpoint.Endpoint
}
