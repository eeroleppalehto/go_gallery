package authservice

import (
	"net/http"

	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	key         = []byte("super-secret-key")
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

func (s *SessionService) Login(c echo.Context, q *models.Queries) error {
	sess, err := store.Get(c.Request(), cookieStore)
	if err != nil {
		c.Response().Status = http.StatusInternalServerError
		return err
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := q.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		c.Response().Status = 401
		return err
	}

	isValidPW := ComparePassowrds([]byte(user.Password), password)
	if !isValidPW {
		c.Response().Status = 401
		return err
	}

	sess.Values["authenticated"] = true
	sess.Values["username"] = username
	sess.Save(c.Request(), c.Response())
	return nil
}

func (s *SessionService) Logout(c echo.Context) error {
	sess, err := store.Get(c.Request(), cookieStore)
	if err != nil {
		return err
	}

	sess.Values["authenticated"] = false
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
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
