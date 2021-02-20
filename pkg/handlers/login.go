package handlers

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
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
		default:
			if !checkCredentials(email, password) {
				w.WriteHeader(http.StatusForbidden)
				data.Error = "Invalid credentials"
				break
			}
			w.WriteHeader(http.StatusOK)
			data.Error = "Success"
		}
	}

	_ = t.Execute(w, data)
}

func checkCredentials(email string, password string) bool {
	// temp user model
	type User struct {
		Email string
		Password  int
	}

	var user User

	query, _ := orm.NewQueryBuilder("mysql")

	query.Select("users.email").
		From("users").
		Where("email = ? AND password = ?").
		Limit(1)

	orm.NewOrm().
		Raw(query.String(), email, password).
		QueryRow(&user)

	return user.Email == email
}
