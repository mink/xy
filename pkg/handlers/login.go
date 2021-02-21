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
	Form LoginFormData
}

// plaintext password will not be passed back to template
type LoginFormData struct {
	Email string
}

type User struct {
	Id int
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

		email := r.FormValue("email")
		password := r.FormValue("password")

		user := User{Email: email, Password: password}

		valid := validation.Validation{}
		valid.Valid(&user)

		data.Form = LoginFormData{Email: user.Email}

		if valid.HasErrors() {
			w.WriteHeader(http.StatusUnprocessableEntity)
			for _, err := range valid.Errors {
				data.Error = err.Message
				break
			}
			_ = t.Execute(w, data)
			return
		}

			w.WriteHeader(http.StatusForbidden)
		if !checkCredentials(&user) {
			data.Error = "Invalid credentials"
		} else {
			w.WriteHeader(http.StatusOK)
			data.Error = "Success"
		}
	}

	_ = t.Execute(w, data)
}

func checkCredentials(user *User) bool {

	var userToCheck User

	query, _ := orm.NewQueryBuilder("mysql")

	query.Select("users.id, users.email").
		From("users").
		Where("email = ? AND password = ?").
		Limit(1)

	orm.NewOrm().
		Raw(query.String(), user.Email, user.Password).
		QueryRow(&userToCheck)

	if userToCheck.Email != user.Email {
		return false
	}

	// set user id
	user.Id = userToCheck.Id

	return true
}
