package main

import (
	"github.com/google/wire"
	"mg-member/sandbox/service"
)

var Set = wire.NewSet(
	service.NewAService,
	wire.Bind(new(service.Caller), new(*service.AService)))

func Init() *service.AService {
	panic(wire.Build(Set))
}
