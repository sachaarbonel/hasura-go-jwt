package main

import (
	"fmt"

	"github.com/Sach97/hasura-go-jwt/builder"
)

func main() {
	hasuraClaims := builder.HasuraClaimsBuilder.
		AddRole("editor").
		AddRole("user").
		DefaultRole("user").
		OrgID("1234").
		Build()
	//fmt.Println(hasuraClaims)
	customClaims := builder.CustomClaimsBuilder.
		Subject("1234").
		ExpiresAt(01234).
		Build(hasuraClaims)
	fmt.Println(customClaims)

}
