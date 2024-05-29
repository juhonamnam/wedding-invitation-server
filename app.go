package main

import (
	"database/sql"
	"net/http"

	"github.com/juhonamnam/wedding-invitation-server/httphandler"
	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	mux.Handle("/api/posts", new(httphandler.PostHandler))

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://card.juhonam.shop", "https://juhonamnam.github.io"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: true,
	})

	handler := corHandler.Handler(mux)

	http.ListenAndServe(":8080", handler)
}
