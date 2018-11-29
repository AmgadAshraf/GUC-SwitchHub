package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

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
	tpl = template.Must(template.ParseGlob("/go/src/app/templates/*.html"))

}

//SignIn handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "SignIn.html", nil)

}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "SignUp.html", nil)

}

//Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	emailValue := r.FormValue("email")
	passwordValue := r.FormValue("password")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT email, userpassword FROM signedup")

	if err != nil {
		log.Fatal(err)
		return
	}
	var (
		emaildb    string
		passworddb string
	)

	for rows.Next() {

		if err := rows.Scan(&emaildb, &passworddb); err != nil {
			//Zero rows
			log.Fatal(err)
			log.Printf("Fadya")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if emaildb == emailValue && passworddb == passwordValue {
			log.Printf("Matched")
			tpl.ExecuteTemplate(w, "Home.html", nil)

		} else {
			log.Printf("Not Matched")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}
func main() {

	http.HandleFunc("/", SignIn)
	http.HandleFunc("/SignUp.html", SignUp)
	http.HandleFunc("/Home", Home)
	http.ListenAndServe(":8080", nil)
}
