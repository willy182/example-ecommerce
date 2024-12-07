package domain

import "time"

// Claim data structure
type Claim struct {
	Subject     string
	PhoneNumber string
}

// AccessToken data structure
type AccessToken struct {
	AccessToken string
	ExpiredAt   time.Time
}

// TokenClaims
type TokenClaims struct {
	ExpiresAt   float64 `json:"exp,omitempty"`
	Jti         string  `json:"jti,omitempty"`
	IssuedAt    float64 `json:"iat,omitempty"`
	Subject     string  `json:"sub,omitempty"`
	PhoneNumber string  `json:"phone_number,omitempty"`
}

// AdditionalClaims
type AdditionalClaims struct {
	PhoneNumber string `json:"phone_number,omitempty"`
}
