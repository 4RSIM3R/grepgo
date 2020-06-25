// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dawuh struct {
	ID     string `json:"id"`
	Dawuh  string `json:"dawuh"`
	Qoil   string `json:"qoil"`
	Author *User  `json:"author"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewDawuh struct {
	Dawuh string `json:"dawuh"`
	Qoil  string `json:"qoil"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Password *string `json:"password"`
}
