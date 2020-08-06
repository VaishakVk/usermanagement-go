package helpers

import "github.com/dgrijalva/jwt-go"

const secret = "mysecret"

func SignToken(payload jwt.Claims) (string, error) {
	signed := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := signed.SignedString([]byte(secret))
	return token, err
}

func ValidateToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return token, err
}
