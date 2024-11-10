package routes

import (
	"database/sql"
	"net/http"

	"github.com/Harisny/listku/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.HelloController())
	server.HandleFunc("/listku", controller.ListkuController(db))
	server.HandleFunc("/listku/add", controller.NewCreateSongController(db))
	server.HandleFunc("/listku/update", controller.NewUpdateSongController(db))
	server.HandleFunc("/listku/delete", controller.DeleteSongController(db))
}