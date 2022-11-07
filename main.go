package main

import (
	"fmt"
	"mg-member/access"
)

func main() {
	fmt.Printf("Hello Marsen Go Member")
	o := access.DB{}
	o.Migrate()
	o.Seed()
	r := &Route{}
	r.Setup()
}
