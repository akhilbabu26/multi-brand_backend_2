package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct{
	UserID uint
	Role string
	jwt.RegisteredClaims //is a predefined struct provided by the JWT library that contains the standard JWT fields (claims)
}

// access and referesh token generetor
func GenerateToken(userID uint, role, sceretKey string, duration time.Duration)(string, error){
	claims := TokenClaims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(sceretKey))
}

// tokens validator
func ValidateToken(tokenStr, sceretKey string)(*TokenClaims, error){
	token, err := jwt.ParseWithClaims(
		tokenStr, // This is the JWT string coming from client
		&TokenClaims{}, // this is for “Parse the JWT payload into THIS struct.”
		func(token *jwt.Token)(interface{}, error){
			return []byte(sceretKey), nil  // it will verify the token with sescretkey and  sotore it in token
		},
	)

	if err != nil{
		return nil, err
	}

	claims := token.Claims.(*TokenClaims)
	return claims, nil
}