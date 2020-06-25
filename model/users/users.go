package users

// User is exported struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
