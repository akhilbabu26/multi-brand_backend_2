package admin

// UPDATE USER REQUEST
type UpdateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// BLOCK / UNBLOCK REQUEST
type BlockUserDTO struct {
	IsBlocked bool `json:"is_blocked"`
}