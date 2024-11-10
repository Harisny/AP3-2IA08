package main

import (
	"net/http"

	"github.com/Harisny/listku/database"
	"github.com/Harisny/listku/routes"
)

func main() {
	db := database.InitDatabase()
	
	server := http.NewServeMux()
	
	// route
	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}