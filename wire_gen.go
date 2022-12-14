// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"mg-member/sandbox/service"
)

// Injectors from wire.go:

func Init() *service.AService {
	aService := service.NewAService()
	return aService
}

// wire.go:

var Set = wire.NewSet(service.NewAService, wire.Bind(new(service.Caller), new(*service.AService)))
