package authservice

import (
	"fmt"
	"net/http"

	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/gorilla/sessions"
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

func (s *SessionService) Login(req *http.Request, w http.ResponseWriter, q *models.Queries) (int, error) {
	sess, err := store.Get(req, cookieStore)
	if err != nil {
		return 500, err
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	user, err := q.GetUserByUsername(req.Context(), username)
	if err != nil {
		return 401, err
	}

	isValidPW := ComparePassowrds([]byte(user.Password), password)
	if !isValidPW {
		return 401, err
	}
	sess.Values["authenticated"] = true
	sess.Values["username"] = username
	err = sess.Save(req, w)
	if err != nil {
		fmt.Println("Session fail: ", err)
		return 500, nil
	}
	return 200, nil
}

func (s *SessionService) Logout(req *http.Request, w http.ResponseWriter) (int, error) {
	sess, err := store.Get(req, cookieStore)
	if err != nil {
		return 500, err
	}

	sess.Values["authenticated"] = false
	sess.Options.MaxAge = -1
	sess.Save(req, w)
	return 200, nil
}

func (s *SessionService) IsAuthenticated(req *http.Request) UserSession {
	sess, _ := store.Get(req, cookieStore)

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
