package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/inject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	// listenPort
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	server := grpc.NewServer()
	reflection.Register(server)
	baseBallApi := inject.InitializeAdAPIServer()
	pb.RegisterBaseBallApiServer(server, baseBallApi)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		err = server.Serve(listenPort)
		if err != nil && err != grpc.ErrServerStopped {
			panic(err)
		}
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("signal received: %s\n", <-sig)
	server.GracefulStop()
}
