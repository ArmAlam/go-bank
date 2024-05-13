package main

import "fmt"

func main() {

	server := NewAPIServer(":3000")

	server.Run()

	fmt.Println("REST API with GO")
}
