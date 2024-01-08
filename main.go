package main

import (
	"github.com/angelicalombas/merge-list-api/routes"
)

// @title Merge lists API
// @description API dedicated to merging two ordered lists
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	routes.HandleRequests()
}
