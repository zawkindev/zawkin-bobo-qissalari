package handlers

import (
	"blog/database"
	"blog/models"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := database.DB.Query("SELECT id, title, date FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Date); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	tmpl.Execute(w, posts)
}
