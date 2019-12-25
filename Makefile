all :
    go run ./users/cmd/main.go
    go run ./users/grpc/grpc.go
    sleep 1
    go run ./news/cmd/main.go
    echo "all started"