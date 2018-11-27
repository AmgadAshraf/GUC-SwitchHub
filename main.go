package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "db"
	dbUser     = "root"
	dbPassword = "secret"
	dbName     = "GUCSwitchHubDB"
)

var tpl *template.Template

func init() {
    var err error
    tpl, err = template.ParseFiles(
        "template/SignUp.html")
    if err != nil {
        panic(err) // handle error
    }
}


func SignUp (w http.ResponseWriter,r*http.Request){
	tpl.ExecuteTemplate(w, "SignUp.html",nil)

}



func main() {

	http.HandleFunc("/", SignUp)
	http.ListenAndServe(":3000",nil)

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)

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
