package main

import (
	"api-rest/database"
	"api-rest/routes"
	"fmt"
)

const (
	version float64 = 1.1
)

func introduction() {
	fmt.Println("Starting server..")
	fmt.Println("System version:", version)
}

func main() {
	introduction()
	database.ConnectionDB()
	routes.HandleRequest()
}
