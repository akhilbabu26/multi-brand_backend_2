// package auth

// import (
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/akhilbabu26/multi-brand_backend_2/config"
// 	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"
// 	"github.com/akhilbabu26/multi-brand_backend_2/utils"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// // signup
// func Signup(c *gin.Context){
// 	var body SignupDTO

// 	if err := c.ShouldBindJSON(&body); err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
// 		return
// 	}

// 	if len(body.Password) < 6 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "password too short"})
// 		return
// 	}

// 	if body.Password != body.CPassword{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "password mismatch"})
// 		return
// 	}

// 	var existing models.User
// 	err := config.DB.Where("email = ?", body.Email).First(&existing).Error // if the user email in our db the err == nil means uaer already exists.
// 	if err == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
// 		return  // .Error = Give me the error result of this DB query and it stores in err.
// 	}

// 	hash, err := bcrypt.GenerateFromPassword(
// 		[]byte(body.Password),
// 		bcrypt.DefaultCost,
// 	)
// 	if err != nil{
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant hash password"})
// 		return
// 	}

// 	role := strings.ToLower(body.Role)
// 	if role == ""{
// 		role = "user"
// 	}
	
// 	otp := utils.GenerateOTP()

// 	// store temporary signup data
// 	pendingUsers[body.Email] = PendingSignup{
// 		Name:      body.Name,
// 		Email:     body.Email,
// 		Password:  string(hash),
// 		Role:      role,
// 		OTP:       otp,
// 		ExpiresAt: time.Now().Add(
// 			time.Minute * time.Duration(config.AppConfig.OTP.ExpiryMinutes),
// 		),
// 	}

// 	// send OTP via email 
// 	err = utils.SendOTPEmail(body.Email, otp)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "failed to send otp email",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "OTP sent to email",
// 	})
// }

// // verify otp
// func VerifyOTP(c *gin.Context) {
// 	var body VerifyOTPDTO

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
// 		return
// 	}

// 	pending, exists := pendingUsers[body.Email]
// 	if !exists {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "no signup request"})
// 		return
// 	}

// 	if time.Now().After(pending.ExpiresAt) {
// 		delete(pendingUsers, body.Email)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "otp expired"})
// 		return
// 	}

// 	if pending.OTP != body.OTP {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong otp"})
// 		return
// 	}

// 	// create real user in DB
// 	user := models.User{
// 		Name:      pending.Name,
// 		Email:     pending.Email,
// 		Password:  pending.Password,
// 		Role:      pending.Role,
// 		IsBlocked: false,
// 	}

// 	if err := config.DB.Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
// 		return
// 	}

// 	// remove temp data
// 	delete(pendingUsers, body.Email)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "signup completed",
// 	})
// }

// //Login
// func Login(c *gin.Context){
// 	var body LoginDTO
// 	var user models.User

// 	if err := c.ShouldBindJSON(&body); err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
// 		return
// 	}

// 	err := config.DB.Where("email = ?", body.Email).
// 	First(&user).Error
// 	if err != nil{
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not signup"})
// 		return 
// 	}

// 	if user.IsBlocked {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "user blocked"})
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword(
// 		[]byte(user.Password),
// 		[]byte(body.Password),
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
// 		return
// 	}

// 	access, err := utils.GenerateToken(
// 		user.ID,
// 		user.Role,
// 		config.AppConfig.JWT.AccessSecretKey,
// 		time.Minute*time.Duration(config.AppConfig.JWT.AccessTTLMinutes),
// 	)
// 	if err != nil{
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant create access token"})
// 		return
// 	}

// 	refresh, err := utils.GenerateToken(
// 		user.ID,
// 		user.Role,
// 		config.AppConfig.JWT.RefreshSecretKey,
// 		time.Hour*time.Duration(config.AppConfig.JWT.RefreshTTLHours),
// 	)
// 	if err != nil{
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant create referesh token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"access_token":  access,
// 		"refresh_token": refresh,
// 		"role":          user.Role,
// 	})
// }

// // token validation
// func RefreshToken(c *gin.Context){
// 	var body RefreshDTO

// 	if err := c.ShouldBindJSON(&body); err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid refresh token"})
// 		return
// 	}

// 	claims, err := utils.ValidateToken(
// 		body.RefreshToken,
// 		config.AppConfig.JWT.RefreshSecretKey,
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh"})
// 		return
// 	}

// 	newAccessToken, err := utils.GenerateToken(
// 		claims.UserID,
// 		claims.Role,
// 		config.AppConfig.JWT.AccessSecretKey,
// 		time.Minute*time.Duration(config.AppConfig.JWT.AccessTTLMinutes),
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "cant create access token",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
// }


package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/akhilbabu26/multi-brand_backend_2/config"
	"github.com/akhilbabu26/multi-brand_backend_2/internal/models"
	"github.com/akhilbabu26/multi-brand_backend_2/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


// AUTH HANDLERS

// SIGNUP
func Signup(c *gin.Context) {
	var body SignupDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	if len(body.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password too short"})
		return
	}

	if body.Password != body.CPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password mismatch"})
		return
	}

	var existing models.User
	err := config.DB.Where("email = ?", body.Email).First(&existing).Error

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant hash password"})
		return
	}

	role := strings.ToLower(body.Role)
	if role != "user" {
		role = "user"
	}

	otp := utils.GenerateOTP()

	// thread-safe write
	mu.Lock()
	pendingUsers[body.Email] = PendingSignup{
		Name:      body.Name,
		Email:     body.Email,
		Password:  string(hash),
		Role:      role,
		OTP:       otp,
		ExpiresAt: time.Now().Add(
			time.Minute * time.Duration(config.AppConfig.OTP.ExpiryMinutes),
		),
	}
	mu.Unlock()

	err = utils.SendOTPEmail(body.Email, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send otp email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent to email",
	})
}

// VERIFY OTP
func VerifyOTP(c *gin.Context) {
	var body VerifyOTPDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// thread-safe read
	mu.RLock()
	pending, exists := pendingUsers[body.Email]
	mu.RUnlock()

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no signup request"})
		return
	}

	if time.Now().After(pending.ExpiresAt) {
		mu.Lock()
		delete(pendingUsers, body.Email)
		mu.Unlock()

		c.JSON(http.StatusBadRequest, gin.H{"error": "otp expired"})
		return
	}

	if pending.OTP != body.OTP {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong otp"})
		return
	}

	user := models.User{
		Name:      pending.Name,
		Email:     pending.Email,
		Password:  pending.Password,
		Role:      pending.Role,
		IsBlocked: false,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	mu.Lock()
	delete(pendingUsers, body.Email)
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "signup completed",
	})
}

// LOGIN
func Login(c *gin.Context) {
	var body LoginDTO
	var user models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := config.DB.Where("email = ?", body.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not signup"})
		return
	}

	if user.IsBlocked {
		c.JSON(http.StatusForbidden, gin.H{"error": "user blocked"})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	access, err := utils.GenerateToken(
		user.ID,
		user.Role,
		config.AppConfig.JWT.AccessSecretKey,
		time.Minute*time.Duration(config.AppConfig.JWT.AccessTTLMinutes),
	)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant create access token"})
		return
	}

	refresh, err := utils.GenerateToken(
		user.ID,
		user.Role,
		config.AppConfig.JWT.RefreshSecretKey,
		time.Hour*time.Duration(config.AppConfig.JWT.RefreshTTLHours),
	)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant create referesh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"role":          user.Role,
	})
}

// REFRESH TOKEN
func RefreshToken(c *gin.Context) {
	var body RefreshDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid refresh token"})
		return
	}

	claims, err := utils.ValidateToken(
		body.RefreshToken,
		config.AppConfig.JWT.RefreshSecretKey,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh"})
		return
	}

	newAccessToken, err := utils.GenerateToken(
		claims.UserID,
		claims.Role,
		config.AppConfig.JWT.AccessSecretKey,
		time.Minute*time.Duration(config.AppConfig.JWT.AccessTTLMinutes),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cant create access token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

// FORGOT PASSWORD
func ForgotPassword(c *gin.Context){
	var body ForgotPasswordDOT

	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input"})
		return
	}

	// check user exists
	var user models.User
	err := config.DB.Where("email = ?", body.Email).First(&user).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// generate OTP
	otp := utils.GenerateOTP()

	// store reset OTP
	resetMu.Lock()
	resetOTPs[body.Email] = PendingReset{
		Email: body.Email,
		OTP: otp,
		ExpiresAt: time.Now().Add(
			time.Minute * time.Duration(config.AppConfig.OTP.ExpiryMinutes),
		),
	}
	resetMu.Unlock()

	// send email
	if err := utils.SendOTPEmail(body.Email, otp); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send otp",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password reset otp sent",
	})
}

// RESET PASSWORD
func ResetPassword(c *gin.Context){
	var body ResetPasswordDOT

	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// read OTP safely
	resetMu.RLock() // Multiple goroutines to read at the same time and cannot modify data while readers are active
	pending, exists := resetOTPs[body.Email]
	resetMu.RUnlock()

	if !exists{
		c.JSON(http.StatusBadRequest, gin.H{"error": "no reset request"})
		return
	}

	// check expiry
	if time.Now().After(pending.ExpiresAt) {
		resetMu.Lock()
		delete(resetOTPs, body.Email)
		resetMu.Unlock()

		c.JSON(http.StatusBadRequest, gin.H{"error": "otp expired"})
		return
	}

	// verify otp
	if pending.OTP != body.OTP{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid otp"})
		return
	}

	// hash new password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(body.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// update db
	err = config.DB.Model(&models.User{}).
		Where("email = ?", body.Email).
		Update("password", string(hash)).Error

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update faild"})
		return
	}

	// delete OTP after success
	resetMu.Lock()
	delete(resetOTPs, body.Email)
	resetMu.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "password reset successful"})
}
