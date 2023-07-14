package common

import (
	"github/adekang/gin-demo/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expiarationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expiarationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
			// 谁发放的token
			Issuer: "adekang",
			// 主题
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
