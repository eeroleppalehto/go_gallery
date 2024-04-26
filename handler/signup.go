package handler

import (
	"github.com/eeroleppalehto/go_gallery/models"
	"github.com/eeroleppalehto/go_gallery/views/signup"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (r *RouteHandler) SignupShow(c echo.Context) error {
	return render(c, signup.Form(signup.NewSignupForm()))
}

func (r *RouteHandler) Signup(c echo.Context) error {
	signupForm := signup.SignupForm{
		Username:        c.FormValue("username"),
		Email:           c.FormValue("email"),
		Password:        c.FormValue("password"),
		ConfirmPassword: c.FormValue("confirm-password"),
		FormError:       signup.None,
	}

	// Check if username exists
	usernameExists, err := r.Queries.UsernameExists(c.Request().Context(), signupForm.Username)
	if err != nil {
		c.Response().Status = 500
		signupForm.FormError = signup.UnknownError
		return render(c, signup.Form(signupForm))
	}
	if usernameExists {
		c.Response().Status = 422
		signupForm.FormError = signup.UsernameExists
		return render(c, signup.Form(signupForm))
	}

	// Check if email exists
	emailExists, err := r.Queries.EmailExists(c.Request().Context(), signupForm.Email)
	if err != nil {
		c.Response().Status = 500
		signupForm.FormError = signup.UnknownError
		return render(c, signup.Form(signupForm))
	}
	if emailExists {
		c.Response().Status = 422
		signupForm.FormError = signup.EmailExists
		return render(c, signup.Form(signupForm))
	}

	// Check if passwords match
	if signupForm.Password != signupForm.ConfirmPassword {
		c.Response().Status = 422
		signupForm.FormError = signup.PasswordsDontMatch
		return render(c, signup.Form(signupForm))
	}

	// TODO: Password validation

	hashBytes, err := HashPassword([]byte(signupForm.Password))
	if err != nil {
		c.Response().Status = 500
		signupForm.FormError = signup.UnknownError
		return render(c, signup.Form(signupForm))
	}
	_, err = r.Queries.CreateUser(c.Request().Context(), models.CreateUserParams{
		Username: signupForm.Username,
		Email:    signupForm.Email,
		Password: string(hashBytes),
	})
	if err != nil {
		c.Response().Status = 500
		signupForm.FormError = signup.UnknownError
		return render(c, signup.Form(signupForm))
	}

	return render(c, signup.Success(signupForm.Username))
}

// TODO: Is this worth?
func HashPassword(pwBytes []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pwBytes, 14)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
