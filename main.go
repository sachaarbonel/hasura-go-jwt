package main

import (
	"fmt"

	"github.com/Sach97/hasura-go-jwt/builder"
)

type HTTPSHasuraIoJwtClaims struct {
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraUserID       string   `json:"x-hasura-user-id"`
	XHasuraOrgID        string   `json:"x-hasura-org-id"`
	XHasuraCustom       string   `json:"x-hasura-custom"`
}

func main() {
	// var dat map[string]interface{}
	// &dat()
	// byt := []byte(`{"https://hasura.io/jwt/claims":6.13}`)
	// if err := json.Unmarshal(byt, &dat); err != nil {
	// 	panic(err)
	// }
	dat := map[string]interface{}{
		"https://hasura.io/jwt/claims": HTTPSHasuraIoJwtClaims{},
	}
	customClaims := builder.CustomClaimsBuilder.
		Subject("1234").
		ExpiresAt(01234).
		Build(dat)
	fmt.Println(customClaims)

}
