package handlers

import (
	"html/template"
	"net/http"
)

type LoginTemplateData struct {
	Error string
}

func Login(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("views/login.html")

	data := LoginTemplateData{}

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == "POST" {

		email := r.FormValue("email")
		password := r.FormValue("password")

		switch {
		case len(email) == 0:
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.Error = "Enter your email address"
			break
		case len(password) == 0:
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.Error = "Enter your password"
			break
		case password != "test":
			w.WriteHeader(http.StatusForbidden)
			data.Error = "Invalid credentials"
			break
		case password == "test":
			w.WriteHeader(http.StatusOK)
			data.Error = "Success"
		}
	}

	_ = t.Execute(w, data)
}
