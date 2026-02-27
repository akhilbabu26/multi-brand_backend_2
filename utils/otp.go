package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP creates a secure 6-digit OTP
func GenerateOTP() string {

	// max = 1000000 (000000 → 999999)
	max := big.NewInt(1000000)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		// fallback (extremely rare)
		return "000000"
	}

	return fmt.Sprintf("%06d", n.Int64())
}