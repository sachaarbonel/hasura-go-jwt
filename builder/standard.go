package builder

import (
	"encoding/base64"

	"github.com/dgrijalva/jwt-go"
	"github.com/lann/builder"
)

type standardClaimsBuilder builder.Builder

func (b standardClaimsBuilder) Subject(subject string) standardClaimsBuilder {
	return builder.Set(b, "Subject", base64.StdEncoding.EncodeToString([]byte(subject))).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) Audience(audience string) standardClaimsBuilder {
	return builder.Set(b, "Audience", audience).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) ExpiresAt(expiresAt int64) standardClaimsBuilder {
	return builder.Set(b, "ExpiresAt", expiresAt).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) ID(ID string) standardClaimsBuilder {
	return builder.Set(b, "ID", ID).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) IssuedAt(issuedAt int64) standardClaimsBuilder {
	return builder.Set(b, "IssuedAt", issuedAt).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) Issuer(issuer string) standardClaimsBuilder {
	return builder.Set(b, "Issuer", issuer).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) NotBefore(notBefore int64) standardClaimsBuilder {
	return builder.Set(b, "NotBefore", notBefore).(standardClaimsBuilder)
}

func (b standardClaimsBuilder) Build() jwt.StandardClaims {
	return builder.GetStruct(b).(jwt.StandardClaims)
}

var StandardClaimsBuilder = builder.Register(standardClaimsBuilder{}, jwt.StandardClaims{}).(standardClaimsBuilder)
