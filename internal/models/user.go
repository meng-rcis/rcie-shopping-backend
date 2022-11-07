package models

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RoleId    uint   `json:"role_id"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
}
