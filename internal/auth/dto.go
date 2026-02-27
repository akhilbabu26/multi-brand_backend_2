// package auth

// type SignupDTO struct{
// 	Name 	string `json:"name"`
// 	Email	string `json:"email"`
// 	Password  string `json:"password"`
// 	CPassword string `json:"cPassword"`
// 	Role	string `json:"role"`
// }

// type LoginDTO struct{
// 	Email string `json:"email"`
// 	Password string `json:"password"`
// }

// type RefreshDTO struct{
// 	RefreshToken string `json:"refresh_token"`
// }

// type VerifyOTPDTO struct {
// 	Email string `json:"email"`
// 	OTP   string `json:"otp"`
// }

package auth

//
// ======================================================
// SIGNUP REQUEST DTO
// ======================================================
//

// SignupDTO represents signup request payload
type SignupDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"cPassword"`
	Role      string `json:"role"`
}

//
// ======================================================
// LOGIN REQUEST DTO
// ======================================================
//

// LoginDTO represents login request payload
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//
// ======================================================
// REFRESH TOKEN REQUEST DTO
// ======================================================
//

// RefreshDTO represents refresh token payload
type RefreshDTO struct {
	RefreshToken string `json:"refresh_token"`
}

//
// ======================================================
// OTP VERIFICATION REQUEST DTO
// ======================================================
//

// VerifyOTPDTO represents OTP verification payload
type VerifyOTPDTO struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}