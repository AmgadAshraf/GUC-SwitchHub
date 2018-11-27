package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=root password=secret dbname=GUCSwitchHubDB sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, fname FROM signedup")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id    string
			fname string
		)
		if err := rows.Scan(&id, &fname); err != nil {
			log.Fatal(err)
		}

		log.Printf("My name is %s, with ID %s", fname, id)
	}

	log.Print("Connected!")
}
