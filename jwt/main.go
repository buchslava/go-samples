package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid uint64, username string) (string, error) {
	var err error
	//this should be in an env file
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["user_name"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	token, err := CreateToken(123, "agent 007")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

	tokenString := token
	claims := jwt.MapClaims{}
	token1, err1 := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	fmt.Println(token1, err1)

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
