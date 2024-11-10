package models

type SignupRes struct {
	ID         string `json:"id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	AuthSource string `json:"auth_source"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type LoginRes struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
