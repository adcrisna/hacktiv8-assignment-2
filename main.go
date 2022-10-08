package main

import (
	"assignment-2/router"
	"fmt"
)

func main() {
	router := router.RouterOrder()
	err := router.Run(":3000")
	if err != nil {
		fmt.Println("Error Saat Connect Ke DB", err)
	}
}
