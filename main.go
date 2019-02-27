package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sach97/hasura-go-jwt/builder"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
	HasuraClaims builder.HasuraClaims `json:"https://hasura.io/jwt/claims"`
}

func main() {
	now := time.Now()
	expires := now.Add(24 * time.Hour * 30)

	hasuraClaims := builder.HasuraClaimsBuilder.
		AddRole("user").
		DefaultRole("user").
		UserID("bhr8su6r85etb5pes9j0").
		Build()
	//fmt.Println(hasuraClaims)

	standardClaims := builder.StandardClaimsBuilder.
		Subject("bhr8su6r85etb5pes9j0").
		ExpiresAt(expires.Unix()).Build()

	customClaims := CustomClaims{
		StandardClaims: standardClaims,
		HasuraClaims:   hasuraClaims,
	}
	b, err := json.Marshal(customClaims)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b), "\n")

	mySigningKey := []byte("e4f05996bffbe1fd0d1dd52dae6fec5dde3bafff4eed4064ba515cfaf3fabee905e36f3a44de129600bab880a21043ef6c21866d21d9440bc41a4325cb29405c")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("%v\n", ss)

	token, err = jwt.ParseWithClaims(ss, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
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
