package auth

import "time"

type PendingSignup struct {
	Name      string
	Email     string
	Password  string
	Role      string
	OTP       string
	ExpiresAt time.Time
}

var pendingUsers = map[string]PendingSignup{}