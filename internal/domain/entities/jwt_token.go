package entities

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Role             string `json:"role"`
	OrganizationCode string `json:"organization_code"`
	jwt.RegisteredClaims
}
