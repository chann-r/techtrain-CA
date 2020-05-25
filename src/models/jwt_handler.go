package models

import "github.com/dgrijalva/jwt-go"

type JwtToken struct {
  Token *jwt.Token
}

// 署名するメソッド
func (token JwtToken) SignedString(key []byte) (string, error) {
	return token.Token.SignedString(key)
}
