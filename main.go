package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/mysql?parseTime=true")
	if err != nil {
		log.Println("There was an error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("There was an error pinging the database: ", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hello")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	err = http.ListenAndServe(":5000", r)

	if err != nil {
		log.Fatal(err)
	}

}
