package session

import (
	"github.com/gorilla/sessions"
	"os"
)

func CreateSessionStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}
