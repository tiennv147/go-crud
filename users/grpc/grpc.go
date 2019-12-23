package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phamtanlong/go-crud/users/pb"
	"github.com/phamtanlong/go-crud/users/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	grpcAddr = ":7770"
)

func main() {
	// init database

	db := initDatabase()
	defer db.Close()

	// listen and serve

	userService := service.UserService{DB:db}

	// grpc
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Printf("> can not create grpc net.Listener")
	}

	authenServer := service.AuthenticationServer{UserService: userService}
	grpcServer := grpc.NewServer()

	pb.RegisterAuthenticationServer(grpcServer, &authenServer)

	log.Printf("> start service grpc")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("> can not call grpcService.Serve")
	}

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


