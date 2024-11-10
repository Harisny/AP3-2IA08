package controller

import "net/http"

func HelloController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halo Dunia"))
	}
}