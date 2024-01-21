package main

import (
	"bialekredki/card-validator/internal"
	"fmt"
)

func main() {
	fmt.Println("Starting server!")
	internal.Serve(8080)
}
