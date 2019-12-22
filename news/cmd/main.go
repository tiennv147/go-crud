package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phamtanlong/go-crud/news/endpoints"
	http1 "github.com/phamtanlong/go-crud/news/http"
	service "github.com/phamtanlong/go-crud/news/service"
	"log"
	"net"
	http "net/http"
)

func main() {
	log.Println("> start service news")

	// init database

	db := initDatabase()
	defer db.Close()

	// listen and serve

	service := service.NewsService{DB:db}

	points := endpoints.Endpoints{
		Create: endpoints.MakeCreateEndpoint(service),
		Update: endpoints.MakeUpdateEndpoint(service),
		Delete: endpoints.MakeDeleteEndpoint(service),
		Read:   endpoints.MakeReadEndpoint(service),
	}

	handler := http1.NewHTTPHandler(points)

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic("> can not create net.Listener")
	}

	http.Serve(listener, handler)
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go-crud?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&service.News{})

	return db
}
