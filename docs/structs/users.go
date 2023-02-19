package structs

type User struct {
	ID       int    `json:"id"`
	PASSWORD string `json:"password"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Role     int    `json:"role"`
}

// 0 - client, 1 - admin, 2 - seller
