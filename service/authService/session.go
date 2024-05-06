package authservice

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	key         = []byte(os.Getenv("SECRET"))
	cookieStore = "session"
	store       = sessions.NewCookieStore(key)
)

type SessionService struct{}

type UserSession struct {
	IsAuthenticated bool
	Username        string
}

func (s *SessionService) Init() {
	store.Options.HttpOnly = true
	store.Options.Secure = true
}

func (s *SessionService) Login(r *http.Request, w http.ResponseWriter, db *sql.DB) int {
	sess, err := GetSession(r)
	if err != nil {
		return http.StatusBadRequest
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	queries := models.New(db)

	user, err := queries.GetUserByUsername(r.Context(), username)
	if err != nil {
		return http.StatusUnauthorized
	}

	isValidPW := ComparePassowrds([]byte(user.Password), password)
	if !isValidPW {
		return http.StatusUnauthorized
	}

	err = SaveSession(r, w, sess, username)
	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	sess, err := store.Get(r, cookieStore)
	if err != nil {
		return nil, err
	}

	return sess, err
}

func SaveSession(r *http.Request, w http.ResponseWriter, sess *sessions.Session, username string) error {
	sess.Values["authenticated"] = true
	sess.Values["username"] = username
	sess.Save(r, w)
	return nil
}

func (s *SessionService) Logout(r *http.Request, w http.ResponseWriter) error {
	sess, err := store.Get(r, cookieStore)
	if err != nil {
		return err
	}

	sess.Values["authenticated"] = false
	sess.Options.MaxAge = -1
	sess.Save(r, w)
	return nil
}

func (s *SessionService) IsAuthenticated(c echo.Context) UserSession {
	sess, _ := store.Get(c.Request(), cookieStore)

	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		return UserSession{
			false,
			"",
		}
	}

	return UserSession{
		sess.Values["authenticated"].(bool),
		sess.Values["username"].(string),
	}
}
