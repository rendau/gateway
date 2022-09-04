package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

func TestJwk(t *testing.T) {
	jwks, err := keyfunc.Get("http://localhost:7575/jwk/set", keyfunc.Options{
		RefreshInterval: 10 * time.Minute,
		RefreshErrorHandler: func(err error) {
			fmt.Println("Jwks refresh error", err)
		},
	})
	require.Nil(t, err)

	jwtToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6Inplb24ubWVjaHRhLm1hcmtldCIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDY4MDIzMjYsImlzcyI6Im1lY2h0YS5tYXJrZXQiLCJyb2xlcyI6WyJtYXRyaXgiXX0.WcBUPywgaqeAReNXnJu61idHhnIM7nLBxkHLaDFnQCFRilk9dATJg4LuGmRgtMUawptZZEAHD3n_roQ1BwyW5gsZETCWkBRfyW98plRHSmAcxzuYKxuTT1MRMMMogPess3BTjjmT6sRANxsK6GkrM0LZajbJaJF5h3wPbp2CCW2WgpHFE5TV0P-8irHVPOxFZOHUELR6kdt1ke7IZ4JDE76-oaomNUOaekTOr2ZWMj3Q4HPbp3y1M_PDIrxiE_39xJaaj_n4J7oYVY9hY_EwFeOOKsxhBGsLdhMNxK8m4f4CyrJt5opqH0mzpzQ6AtN5XO93VC_-nU9oPPNBMoZICQ"

	token, _ := jwt.Parse(jwtToken, jwks.Keyfunc)
	require.False(t, token.Valid)
}
