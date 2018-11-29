package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/GeertJohan/go.rice"
   "github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "db"
	dbUser     = "root"
	dbPassword = "secret"
	dbName     = "GUCSwitchHubDB"
)

//SignIn handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"/go/src/app/SignIn.html")
	if err != nil {
		panic(err) // handle error
	}
	tpl.ExecuteTemplate(w, "SignIn.html", nil)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"/go/src/app/SignUp.html")
	if err != nil {
		panic(err) // handle error
	}
	tpl.ExecuteTemplate(w, "SignUp.html", nil)

}

func main() {

	router := mux.NewRouter()
    router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("templates").HTTPBox()))
	//log.Fatal(http.ListenAndServe(":8080", router))
	
	
	
	
	//router.HandleFunc("/", SignIn)
	//http.HandleFunc("/", SignIn)
	http.ListenAndServe(":8080", router)

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
