package users

type Employee struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
