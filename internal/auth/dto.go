package auth

type SignupDTO struct{
	Name 	string `json:"name"`
	Email	string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"cPassword"`
	Role	string `json:"role"`
}

type LoginDTO struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type RefreshDTO struct{
	RefreshToken string `json:"refresh_token"`
}