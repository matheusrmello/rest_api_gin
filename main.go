package main

import (
	"github.com/matheusrmello/api-go-gin/database"
	"github.com/matheusrmello/api-go-gin/routes"
)

func main() {
	database.ConnectionDB()
	routes.HandleRequests()
}
