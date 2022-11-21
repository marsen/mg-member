package models

import "mg-member/sandbox/service"

type Caller interface {
	Call(i service.AInput) (o service.AOutput)
}
