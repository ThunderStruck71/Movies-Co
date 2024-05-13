package main

import (
	"back"
	"fmt"
	"log"
)

func main() {
	port := "8080"
	srv := new(back.Server)

	fmt.Printf("Server was running on port %s...\n", port)
	if err := srv.Run(port); err != nil {
		log.Fatalf("Something wrong while running server: %s", err.Error())
	}
}
