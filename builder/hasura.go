package builder

import (
	"encoding/base64"

	"github.com/lann/builder"
)

type HasuraClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	UserID       string   `json:"x-hasura-user-id"`
	OrgID        string   `json:"x-hasura-org-id"`
	Custom       string   `json:"x-hasura-custom"`
}

type hasuraClaimsBuilder builder.Builder

func (b hasuraClaimsBuilder) AddRole(allowedRole string) hasuraClaimsBuilder {
	return builder.Append(b, "AllowedRoles", allowedRole).(hasuraClaimsBuilder)
}

func (b hasuraClaimsBuilder) DefaultRole(defaultRole string) hasuraClaimsBuilder {
	return builder.Set(b, "DefaultRole", defaultRole).(hasuraClaimsBuilder)
}

func (b hasuraClaimsBuilder) UserID(userID string) hasuraClaimsBuilder {
	return builder.Set(b, "UserID", userID).(hasuraClaimsBuilder)
}

func (b hasuraClaimsBuilder) OrgID(orgID string) hasuraClaimsBuilder {
	return builder.Set(b, "OrgID", base64.StdEncoding.EncodeToString([]byte(orgID))).(hasuraClaimsBuilder)
}

func (b hasuraClaimsBuilder) Custom(custom string) hasuraClaimsBuilder {
	return builder.Set(b, "Custom", custom).(hasuraClaimsBuilder)
}

func (b hasuraClaimsBuilder) Build() HasuraClaims { //Get
	return builder.GetStruct(b).(HasuraClaims)
}

var HasuraClaimsBuilder = builder.Register(hasuraClaimsBuilder{}, HasuraClaims{}).(hasuraClaimsBuilder)
