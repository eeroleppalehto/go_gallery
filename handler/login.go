package handler

import (
	"net/http"
	"os"
	"time"

	authservice "github.com/eeroleppalehto/go_gallery/service/authService"
	"github.com/eeroleppalehto/go_gallery/views/login"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

func (r *RouteHandler) LoginForm(c echo.Context) error {
	return render(c, login.LoginForm())
}

func (r *RouteHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := r.Queries.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		c.Response().Status = 401
		return render(c, login.LoginForm())
	}

	as := authservice.AuthService{}

	isValidPW := as.ComparePassowrds([]byte(user.Password), password)
	if !isValidPW {
		c.Response().Status = 401
		return render(c, login.LoginForm())
	}

	claims := &authservice.JwtCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)

	cookie.HttpOnly = true
	cookie.Secure = true

	cookie.Name = "token"
	cookie.Value = t
	// cookie.Value = fmt.Sprintf("Bearer %s", t)
	cookie.Expires = time.Now().Add(time.Hour * 48)

	c.SetCookie(cookie)

	return render(c, login.Success())
}

func (r *RouteHandler) Logout(c echo.Context) error {
	t, err := c.Cookie("token")
	if err != nil {
		return render(c, login.Success())
	}

	t.Expires = time.Now().Add(-time.Hour)

	c.SetCookie(t)

	return render(c, login.Success())
}
