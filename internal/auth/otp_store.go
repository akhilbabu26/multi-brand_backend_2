package auth

import (
	"sync"
	"time"
)

// SIGNUP OTP STORE: PendingSignup stores temporary signup data, until OTP verification is completed.
type PendingSignup struct {
	Name      string
	Email     string
	Password  string
	Role      string
	OTP       string
	ExpiresAt time.Time
}

// IN-MEMORY OTP STORE: pendingUsers stores temporary users waiting for OTP verification
var pendingUsers = map[string]PendingSignup{}

// CONCURRENCY CONTROL: mu protects pendingUsers from concurrent access
var mu sync.RWMutex // It allows Multiple readers at the same time,Only one writer at a time,No readers while writing

// SIGNUP OTP STORE
// PendingReset stores reset password OTP data
type PendingReset struct{
	Email string
	OTP string
	ExpiresAt time.Time
}

// reset password OTP map
var resetOTPs = map[string]PendingReset{}

// lock for reset OTP
var resetMu sync.RWMutex