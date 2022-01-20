package tools

import (
	"Task3/config"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	AuthenticationToken = iota
	RefreshToken
)

type (
	tokenClaims struct {
		Userid int `json:"userid"`
		Type   int `json:"type"`
		jwt.StandardClaims
	}
)

func JWTGenerate(id int, issuer string, type_ int) (token string) {
	claims := tokenClaims{
		id,
		type_,
		jwt.StandardClaims{
			Issuer: issuer,
		},
	}
	if type_ == AuthenticationToken {
		claims.ExpiresAt = time.Now().Add(config.AuthTokenExpireDuration).Unix()
	} else if type_ == RefreshToken {
		claims.ExpiresAt = time.Now().Add(config.RefreshTokenExpireDuration).Unix()
	} else {
		panic("type not found")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = t.SignedString(config.TokenSecret)
	return
}

func JWTVerify(signature string) (int, int, error) {
	var claims tokenClaims
	_, err := jwt.ParseWithClaims(signature, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.TokenSecret, nil
	})
	if err != nil {
		return 0, 0, err
	}
	return claims.Userid, claims.Type, nil
}
