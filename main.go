package main

import jwt "github.com/dgrijalva/jwt-go"
import "log"

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
	jwt.StandardClaims
}

func createTokenString() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Name: "konojunya",
		Age:  20,
	})

	tokenstring, err := token.SignedString([]byte("secretstring"))
	if err != nil {
		log.Fatal(err)
	}

	return tokenstring
}

func main() {
	tokenstring := createTokenString()
	log.Println(tokenstring)

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretstring"), nil
	})

	log.Println(token.Claims, err)
}
