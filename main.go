package main

import (
	"fmt"
	"mg-member/access"
)

func main() {
	fmt.Printf("Hello Marsen Go Member")
	o := access.DB{}
	o.Migrate()
	r := &Route{}
	r.Setup()
}
