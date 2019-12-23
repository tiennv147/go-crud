package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phamtanlong/go-crud/users/endpoints"
	"github.com/phamtanlong/go-crud/users/http"
	"github.com/phamtanlong/go-crud/users/service"
	"log"
	"net"
	http1 "net/http"
)

const (
	httpAddr = ":7777"
)

func main() {
	// init database

	db := initDatabase()
	defer db.Close()

	// listen and serve

	service := service.UserService{DB:db}

	points := endpoints.Endpoints{
		Register: endpoints.MakeRegisterEndpoint(service),
		Login:    endpoints.MakeLoginEndpoint(service),
		Verify:   endpoints.MakeVerifyEndpoint(service),
	}

	listener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		log.Fatalf("> can not create net.Listener %v", err)
	}

	handler := http.NewHTTPHandler(points)

	log.Println("> start service auth")
	http1.Serve(listener, handler)
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go-crud?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&service.User{})

	return db
}


