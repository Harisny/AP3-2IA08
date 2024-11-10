package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Musik struct {
	Id     string
	Title  string
	Artist string
	Genre  string
}

func ListkuController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, artist, genre  FROM musik")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var songs []Musik
		for rows.Next() {
			var song Musik

			err = rows.Scan(
				&song.Id,
				&song.Title,
				&song.Artist,
				&song.Genre,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			songs = append(songs, song)
		}


		fp := filepath.Join("views", "listku.html")
		tmpl, err :=template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["songs"] = songs

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}