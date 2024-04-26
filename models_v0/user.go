package models

type Userhandler struct{}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (h *Userhandler) GetUsers() []User {
	users := []User{
		{ID: 1, Username: "Alexios", Email: "alexx@gmail.com", Password: "Q2werty"},
		{ID: 2, Username: "MoonMoon", Email: "kuunlapsi@luukku.fi", Password: "Q2werty"},
		{ID: 3, Username: "Lumis", Email: "lumis@yahoo.com", Password: "Q2werty"},
		{ID: 4, Username: "TupperPupper", Email: "pupper123@gmail.com", Password: "Q2werty"},
	}

	return users
}
