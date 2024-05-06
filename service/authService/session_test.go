package authservice_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"testing"

	authservice "github.com/eeroleppalehto/go_gallery/service/authService"
)

func TestLogin(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"Correct credentials", `username=EL&password=Q2werty`, 200},
		{"Incorrect credentials", `username=none&password=letters`, 401},
	}

	db := SetupDB(t)
	defer db.Close()

	s := authservice.SessionService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := bytes.NewReader([]byte(tt.input))

			r := httptest.NewRequest(http.MethodPost, "/", creds)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()

			status := s.Login(r, w, db)

			if tt.want != status {
				t.Fatalf("Expected status %d, got %d", tt.want, status)
			}
		})
	}
}

func TestLogut(t *testing.T) {
	var tests = []struct {
		name     string
		request  *http.Request
		response *httptest.ResponseRecorder
		want     bool
	}{
		{
			"Valid session cookie",
			SetupSession(t, "user"),
			httptest.NewRecorder(),
			false,
		},
		{
			"Invalid session cookie",
			SetupInavalidSession(t),
			httptest.NewRecorder(),
			true,
		},
	}

	s := authservice.SessionService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Logout(tt.request, tt.response)
			isError := err == nil
			if tt.want == isError {
				t.Fatalf("Expected %t, got %t", tt.want, isError)
			}
		})
	}
}

func SetupDB(t testing.TB) *sql.DB {
	LoadEnv(t)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gollery?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD")))
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}

	return db
}

func LoadEnv(t testing.TB) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error while loading variables from .env: %s", err)
	}
}

func SetupSession(t testing.TB, username string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	sess, _ := authservice.GetSession(r)
	authservice.SaveSession(r, w, sess, username)

	return r
}

func SetupInavalidSession(t testing.TB) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/", nil)

	cookie := &http.Cookie{
		Name:  "session",
		Value: "invalid_session_token",
	}

	r.AddCookie(cookie)

	return r
}
