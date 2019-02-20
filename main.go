package main

import (
	"encoding/base64"
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
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: 1516239022,
			Issuer:    "test",
		},
		HTTPSHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user", "editor"},
			XHasuraDefaultRole:  "user",
			XHasuraOrgID:        base64.StdEncoding.EncodeToString([]byte("1234567890")),
			XHasuraCustom:       "custom-value",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("%v\n", ss)

	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyMzkwMjIsImlzcyI6InRlc3QiLCJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWFsbG93ZWQtcm9sZXMiOlsidXNlciIsImVkaXRvciJdLCJ4LWhhc3VyYS1kZWZhdWx0LXJvbGUiOiJ1c2VyIiwieC1oYXN1cmEtdXNlci1pZCI6IiIsIngtaGFzdXJhLW9yZy1pZCI6Ik1USXpORFUyTnpnNU1BPT0iLCJ4LWhhc3VyYS1jdXN0b20iOiJjdXN0b20tdmFsdWUifX0.5etS2jE6D4G4n92UKt4mGwuT1KiPBOfJl1Jgso3MxsQ"
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v\n%v\n", claims.HTTPSHasuraIoJwtClaims, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	})

}
