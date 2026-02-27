// package auth

// import (
// 	"sync"
// 	"time"
// )

// type PendingSignup struct {
// 	Name      string
// 	Email     string
// 	Password  string
// 	Role      string
// 	OTP       string
// 	ExpiresAt time.Time
// }

// var (
// 	pendingUsers = map[string]PendingSignup{}
// 	mu           sync.RWMutex
// )

package auth

import (
	"sync"
	"time"
)

//
// ======================================================
// PENDING SIGNUP MODEL
// ======================================================
//

// PendingSignup stores temporary signup data
// until OTP verification is completed.
type PendingSignup struct {
	Name      string
	Email     string
	Password  string
	Role      string
	OTP       string
	ExpiresAt time.Time
}

//
// ======================================================
// IN-MEMORY OTP STORE
// ======================================================
//

// pendingUsers stores temporary users waiting for OTP verification
var pendingUsers = map[string]PendingSignup{}

//
// ======================================================
// CONCURRENCY CONTROL
// ======================================================
//

// mu protects pendingUsers from concurrent access
var mu sync.RWMutex