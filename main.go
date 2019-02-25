package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

type HTTPSHasuraIoJwtClaims struct {
	XHasuraRole         string   `json:"x-hasura-role"`
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraUserID       string   `json:"x-hasura-user-id"`
	XHasuraOrgID        string   `json:"x-hasura-org-id"`
	XHasuraCustom       string   `json:"x-hasura-custom"`
}

type MyCustomClaims struct {
	jwt.StandardClaims
	HTTPSHasuraIoJwtClaims HTTPSHasuraIoJwtClaims `json:"https://hasura.io/jwt/claims"`
}

func main() {
	mySigningKey := []byte("e4f05996bffbe1fd0d1dd52dae6fec5dde3bafff4eed4064ba515cfaf3fabee905e36f3a44de129600bab880a21043ef6c21866d21d9440bc41a4325cb29405c")
	now := time.Now()
	expires := now.Add(24 * time.Hour * 30)

	// Create the Claims
	claims := MyCustomClaims{
		jwt.StandardClaims{
			Subject:   "bhq03umr85euc4uua4j0",
			ExpiresAt: expires.Unix(),
		},
		HTTPSHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user"},
			XHasuraRole:         "user",
			XHasuraDefaultRole:  "user",
			XHasuraUserID:       "bhq03umr85euc4uua4j0",
			XHasuraCustom:       "custom-value",
		},
	}
	b, err := json.Marshal(claims)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b), "\n")

	//fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("%v\n", ss)

	token, err = jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		b, err := json.Marshal(claims)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(string(b))
		fmt.Printf("%v\n%v\n", claims, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

}
