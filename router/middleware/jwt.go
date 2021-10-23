package middleware

import (
	"errors"
	"github.com/XGHXT/SYOJ-Backend/config"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var ErrorInvalidToekn = errors.New("invalid token")

const TokenExpireDuration = time.Minute * 60 * 6

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt 包

type MyClaims struct {
	UserName string `json:"username"`
	UserID   int64  `json:"user_id"`
	jwt.StandardClaims
}

// 发放 token
func ReleaseToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims := &MyClaims{
		UserName: user.Username,
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "SYOJ",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecret))
	return tokenString, err
}

// 解析 token
func ParseToekn(tokenString string) (*MyClaims, error) {
	var claims = new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, ErrorInvalidToekn
}
