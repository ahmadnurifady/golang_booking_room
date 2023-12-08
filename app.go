package main

import (
	"time"

	"final-project/delivery"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	delivery.NewServer().Run()
	// fmt.Println(token())
}

func token() string {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	return ss
}
