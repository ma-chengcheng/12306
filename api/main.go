package main

import (
	"github.com/mamachengcheng/12306/api/router"
	"log"
)

func main() {
	r := router.InitRouter()

	// Run service
	if err := r.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
