package handlers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"xy/pkg/session"
)

type LoginTemplateData struct {
	Error string
	Form LoginFormData
	Session SessionData
}

type SessionData struct {
	Authenticated bool
	UserId int
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

var store *sessions.CookieStore = session.CreateSessionStore()

func Login(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("views/login.html")

	data := LoginTemplateData{}

	session, _ := store.Get(r, "session")

	if session.Values["authenticated"] == nil {
		session.Values["authenticated"] = false
	}

	if session.Values["user_id"] == nil {
		session.Values["user_id"] = 0
	}

	data.Session = SessionData{}
	data.Session.Authenticated = session.Values["authenticated"].(bool)
	data.Session.UserId = session.Values["user_id"].(int)

	if r.Method == "GET" {
		session.Save(r, w)
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
			for _, err := range valid.Errors {
				data.Error = err.Message
				break
			}
			session.Save(r, w)
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = t.Execute(w, data)
			return
		}

		if !checkCredentials(&user) {
			data.Error = "Invalid credentials"
			session.Save(r, w)
			w.WriteHeader(http.StatusForbidden)

		} else {
			data.Error = "Success"

			// store user to session
			session.Values["authenticated"] = true
			session.Values["user_id"] = user.Id
			session.Save(r, w)
			w.WriteHeader(http.StatusOK)
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
