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

//SignUpLoader handler
func SignUpLoader(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "SignUp.html", nil)
}

//SignUp handler
func SignUp(w http.ResponseWriter, r *http.Request) {

	fnameValues := r.FormValue("firstname")
	lnameValues := r.FormValue("lastname")
	gucidValues := r.FormValue("GUCID")
	emailValues := r.FormValue("signupemail")
	passwordValues := r.FormValue("signuppassword")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}
	var returnedID string
	err = db.QueryRow(`SELECT id FROM signedup WHERE id=$1;`, gucidValues).Scan(&returnedID)
	if err != nil {
		if err == sql.ErrNoRows {

		} else {
			http.Redirect(w, r, "SignUpLoader", http.StatusSeeOther)
			return
		}

	} else {
		http.Redirect(w, r, "SignUpLoader", http.StatusSeeOther)
		return
	}
	var returnedEmail string
	err = db.QueryRow(`SELECT email FROM signedup WHERE email=$1;`, emailValues).Scan(&returnedEmail)
	if err != nil {
		if err == sql.ErrNoRows {

		} else {
			http.Redirect(w, r, "SignUpLoader", http.StatusSeeOther)
			return
		}

	} else {
		http.Redirect(w, r, "SignUpLoader", http.StatusSeeOther)
		return
	}

	sqlStatement := `INSERT into signedup(fname, lname, id, email, userpassword)
	VALUES($1,$2,$3,$4,$5)`
	_, err = db.Exec(sqlStatement, fnameValues, lnameValues, gucidValues, emailValues, passwordValues)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/Home", http.StatusSeeOther)
	http.StatusText(200)

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
	var returnedEmail string
	var returnedPassword string
	err = db.QueryRow(`SELECT email, userpassword FROM signedup WHERE email=$1;`, emailValue).Scan(&returnedEmail, &returnedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "You are not registred. Please Sign Up", http.StatusForbidden)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		log.Fatal(err)
		return
	}

	if emailValue == "" || passwordValue == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if returnedEmail == emailValue && returnedPassword != passwordValue {
		//http.Error(w, "Incorrect Password", http.StatusForbidden)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if returnedEmail == emailValue && returnedPassword == passwordValue {
		tpl.ExecuteTemplate(w, "Home.html", nil)
		http.StatusText(200)
		return
	}

	//http.Redirect(w, r, "/", http.StatusSeeOther)

}
func main() {

	http.HandleFunc("/", SignIn)
	http.HandleFunc("/SignUp", SignUp)
	http.HandleFunc("/SignUpLoader", SignUpLoader)
	http.HandleFunc("/Home", Home)
	http.ListenAndServe(":8080", nil)
}
