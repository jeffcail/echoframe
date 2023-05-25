package jwt

import (
	"errors"
	"time"

	"github.com/echoframe/conf"
	"github.com/golang-jwt/jwt"
)

var (
	expire int64
	secret string
)

func init() {
	expire = conf.Config.Jwt.EXPIRE
	secret = conf.Config.Jwt.SECRET
}

type JwtClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken
func GenerateToken(claims *JwtClaims) (string, error) {
	c := JwtClaims{
		ID:       claims.ID,
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expire)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

// ParseToken
func ParseToken(ts string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(ts, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
