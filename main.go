package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/gomail.v2"
)

//Config JSON
type Config struct {
	DbHost     string `json:"dbHost"`
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
	DbName     string `json:"dbName"`
}

var dbHost string
var dbUser string
var dbPassword string
var dbName string

var sessionEmail string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/go/src/app/templates/*.html"))

	var configuration Config
	file, err := os.Open("/go/src/app/config.json")
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&configuration)

	dbHost = configuration.DbHost
	dbUser = configuration.DbUser
	dbPassword = configuration.DbPassword
	dbName = configuration.DbName
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
			http.Redirect(w, r, "/SignIn.html", http.StatusSeeOther)
			return
		}
		log.Fatal(err)
		return
	}

	if emailValue == "" || passwordValue == "" {
		http.Redirect(w, r, "/SignIn.html", http.StatusSeeOther)
		return
	}

	if returnedEmail == emailValue && returnedPassword != passwordValue {
		//http.Error(w, "Incorrect Password", http.StatusForbidden)
		http.Redirect(w, r, "/wrongPassword.html", http.StatusSeeOther)
		return
	}

	if returnedEmail == emailValue && returnedPassword == passwordValue {
		sessionEmail = emailValue
		tpl.ExecuteTemplate(w, "Home.html", nil)
		http.StatusText(200)
		return
	}

	//http.Redirect(w, r, "/", http.StatusSeeOther)

}

//SendSwitchRequest handler
func SendSwitchRequest(w http.ResponseWriter, r *http.Request) {

	majorValue := r.FormValue("major")
	currentTutValue := r.FormValue("currentTut")
	desiredTutValue := r.FormValue("desiredTut")
	englishLevelValue := r.FormValue("englishlevel")
	germanLevelValue := r.FormValue("germanlevel")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}
	var returnedUserID string
	err = db.QueryRow(`SELECT id FROM signedup WHERE email=$1;`, sessionEmail).Scan(&returnedUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			//http.Error(w, "You are not registred. Please Sign Up", http.StatusForbidden)
			return
		}
		log.Fatal(err)
		return
	}
	var recievedEmail string
	err = db.QueryRow(`SELECT email FROM switching WHERE email=$1;`, sessionEmail).Scan(&recievedEmail)
	if err != nil {
		if err == sql.ErrNoRows {

		} else {
			http.Redirect(w, r, "/Warning", http.StatusSeeOther)
			return
		}

	} else {
		http.Redirect(w, r, "/Warning", http.StatusSeeOther)
		return
	}

	sqlStatement := `INSERT into switching(id, major, tutorialfrom, tutorialto, germanlevel,englishlevel,email)
	VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err = db.Exec(sqlStatement, returnedUserID, majorValue, currentTutValue, desiredTutValue, germanLevelValue, englishLevelValue, sessionEmail)
	if err != nil {
		http.Redirect(w, r, "/Warning", http.StatusSeeOther)
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT * from switching")
	if err != nil {
		log.Fatal(err)
	}

	var (
		idCmpr        string
		majorCmpr     string
		tfCmpr        string
		ttCmpr        string
		germanCmpr    string
		englishCmpr   string
		emailCmpr     string
		didSwitchCmpr bool
	)
	for rows.Next() {

		if err := rows.Scan(&idCmpr, &majorCmpr, &tfCmpr, &ttCmpr, &germanCmpr, &englishCmpr, &emailCmpr, &didSwitchCmpr); err != nil {
			log.Fatal(err)
		}

		dashValue := returnedUserID[0:2]
		dashCompare := idCmpr[0:2]

		var didSwitchCurrent bool
		err = db.QueryRow(`SELECT didswitch FROM switching WHERE email=$1;`, sessionEmail).Scan(&didSwitchCurrent)
		if err != nil {
			if err == sql.ErrNoRows {
				//http.Error(w, "You are not registred. Please Sign Up", http.StatusForbidden)
				return
			}
			log.Fatal(err)
			return
		}

		if dashValue == dashCompare && majorValue == majorCmpr &&
			germanLevelValue == germanCmpr && englishLevelValue == englishCmpr &&
			didSwitchCmpr == false && didSwitchCurrent == false && currentTutValue == ttCmpr && desiredTutValue == tfCmpr {

			sqlStatement1 := `UPDATE switching SET tutorialfrom = $1, tutorialto = $2, didswitch = $3 WHERE email = $4;`

			_, err = db.Exec(sqlStatement1, "SWITCHED", "SWITCHED", true, sessionEmail)
			if err != nil {
				panic(err)
			}

			_, err = db.Exec(sqlStatement1, "SWITCHED", "SWITCHED", true, emailCmpr)
			if err != nil {
				panic(err)
			}

			m := gomail.NewMessage()
			m.SetHeader("From", "gucswitchhub@gmail.com")
			m.SetHeader("To", sessionEmail, emailCmpr)
			m.SetHeader("Subject", "Congratulations!")
			m.SetBody("text/plain", "Dear GUC Switch Hub user, Congratulations! Your tutorial group have been successfully switched!\n --This is an autogenerated email from the GUC Switch Hub Web App.")

			d := gomail.NewDialer("smtp.gmail.com", 587, "gucswitchhub@gmail.com", "GucSwitchHub2018")

			// Send the email to Bob, Cora and Dan.
			if err := d.DialAndSend(m); err != nil {
				panic(err)
			}

			log.Printf("SENT!")

		} else {

		}

	}

	rows1, err := db.Query("SELECT * from switching")
	if err != nil {
		log.Fatal(err)

	}

	var (
		id    string
		major string
		tf    string
		tt    string
		g     string
		e     string
		email string
		did   bool
	)
	for rows1.Next() {

		if err := rows1.Scan(&id, &major, &tf, &tt, &g, &e, &email, &did); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %s Major: %s From: %s To: %s German: %s English: %s Email: %s DidSwitch: %t", id, major, tf, tt, g, e, email, did)

	}
	log.Printf("FASEL")

	http.Redirect(w, r, "/Final", http.StatusSeeOther)
	return

}

//Final Handler
func Final(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Final.html", nil)
}

//Warning Handler
func Warning(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Warning.html", nil)
}

func main() {

	http.Handle("/", http.FileServer(assetFS()))

	//http.HandleFunc("/", SignIn)
	http.HandleFunc("/SignUp", SignUp)
	http.HandleFunc("/SignUpLoader", SignUpLoader)
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/SendSwitchRequest", SendSwitchRequest)
	http.HandleFunc("/Final", Final)
	http.HandleFunc("/Warning", Warning)
	http.ListenAndServe(":8080", nil)

}
