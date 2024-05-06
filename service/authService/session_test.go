package authservice_test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"

	_ "github.com/go-sql-driver/mysql"

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds := bytes.NewReader([]byte(tt.input))

			r := httptest.NewRequest(http.MethodPost, "/", creds)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()

			db, err := sql.Open("mysql", "root:Q2werty@/gollery?parseTime=true")
			if err != nil {
				t.Fatalf("Failed to connect to database: %s", err)
			}

			s := authservice.SessionService{}

			status := s.Login(r.Context(), r, w, db)

			if tt.want != status {
				t.Fatalf("Expected status %d, got %d", tt.want, status)
			}
		})
	}

}
