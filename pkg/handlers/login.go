package handlers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type LoginTemplateData struct {
	Error string
}

type User struct {
	Email string `valid:"Email; MaxSize(100)"`
	Password  string `valid:"Required"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("views/login.html")

	data := LoginTemplateData{}

	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == "POST" {

		user := User{Email: r.FormValue("email"), Password: r.FormValue("password")}

		valid := validation.Validation{}
		valid.Valid(&user)

		if valid.HasErrors() {
			w.WriteHeader(http.StatusUnprocessableEntity)
			for _, err := range valid.Errors {
				data.Error = err.Message
				break
			}
			_ = t.Execute(w, data)
			return
		}

		if !checkCredentials(user) {
			w.WriteHeader(http.StatusForbidden)
			data.Error = "Invalid credentials"
		} else {
			w.WriteHeader(http.StatusOK)
			data.Error = "Success"
		}
	}

	_ = t.Execute(w, data)
}

func checkCredentials(user User) bool {

	query, _ := orm.NewQueryBuilder("mysql")

	var userToCheck User

	query.Select("users.email").
		From("users").
		Where("email = ? AND password = ?").
		Limit(1)

	orm.NewOrm().
		Raw(query.String(), user.Email, user.Email).
		QueryRow(&userToCheck)

	return userToCheck.Email == user.Email
}
