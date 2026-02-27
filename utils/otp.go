package utils

import (
	"fmt"
	"math/rand" // used to generate pseudo-random numbers
)

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))// Generate a random integer from 0 to 999999 
}