package users

type User struct {
	ID       int    `json:"id"`
	PASSWORD string `json:"password"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
}

type Users struct {
	Users []User `json:"user"`
}
