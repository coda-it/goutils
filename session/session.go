package session

import (
	"github.com/coda-it/goutils/hash"
	"net/http"
	"time"
)

// CreateSessionID - creates a new session ID
func CreateSessionID(user string, pass string, time string) string {
	return hash.String(user + pass + time)
}

// GetSessionID - get user session ID
func GetSessionID(r *http.Request, sessionKey string) (string, error) {
	sessionCookie, err := r.Cookie(sessionKey)

	if err != nil {
		return "", err
	}

	return sessionCookie.Value, nil
}

// ClearSession - remove session cookie
func ClearSession(w http.ResponseWriter, sessionKey string) {
	cookie := http.Cookie{
		Path:    "/",
		Name:    sessionKey,
		Expires: time.Now().Add(-100 * time.Hour),
		MaxAge:  -1}
	http.SetCookie(w, &cookie)
}
