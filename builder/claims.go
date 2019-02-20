package builder

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lann/builder"
)

type CustomClaims struct {
	StandardClaims jwt.StandardClaims
	CustomClaimsI  map[string]interface{}
}

type customClaimsBuilder builder.Builder

func (b customClaimsBuilder) Subject(subject string) customClaimsBuilder {
	return builder.Set(b, "Subject", subject).(customClaimsBuilder)
}

func (b customClaimsBuilder) Audience(audience string) customClaimsBuilder {
	return builder.Set(b, "Audience", audience).(customClaimsBuilder)
}

func (b customClaimsBuilder) ExpiresAt(expiresAt int64) customClaimsBuilder {
	return builder.Set(b, "ExpiresAt", expiresAt).(customClaimsBuilder)
}

func (b customClaimsBuilder) ID(ID string) customClaimsBuilder {
	return builder.Set(b, "ID", ID).(customClaimsBuilder)
}

func (b customClaimsBuilder) IssuedAt(issuedAt int64) customClaimsBuilder {
	return builder.Set(b, "IssuedAt", issuedAt).(customClaimsBuilder)
}

func (b customClaimsBuilder) Issuer(issuer string) customClaimsBuilder {
	return builder.Set(b, "Issuer", issuer).(customClaimsBuilder)
}

func (b customClaimsBuilder) NotBefore(notBefore int64) customClaimsBuilder {
	return builder.Set(b, "NotBefore", notBefore).(customClaimsBuilder)
}

func (b customClaimsBuilder) Build(jsonKey string) CustomClaims { //Get
	standardClaims := builder.GetStruct(b).(jwt.StandardClaims)
	customClaimsI := map[string]interface{}{
		jsonKey: HasuraClaims{},
	}
	return CustomClaims{
		StandardClaims: standardClaims,
		CustomClaimsI:  customClaimsI,
	}
}

var CustomClaimsBuilder = builder.Register(customClaimsBuilder{}, jwt.StandardClaims{}).(customClaimsBuilder)
