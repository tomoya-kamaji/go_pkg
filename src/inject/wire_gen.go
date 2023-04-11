// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/api"
)

// Injectors from wire.go:

func InitializeAdAPIServer() protobuf.BaseBallApiServer {
	baseBallApiServer := api.NewApiServer()
	return baseBallApiServer
}
