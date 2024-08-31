package main

import (
	"blog/database"
	"blog/handlers"
	"net/http"
)

func main() {
	database.Init()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Home)

	http.ListenAndServe(":8080", nil)
}
