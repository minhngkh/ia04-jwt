package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type UserInfo struct {
	Email string `json:"email"`
}

type jwtClaims struct {
	UserInfo
	jwt.RegisteredClaims
}

type JwtToken struct {
	Value     string
	ExpiresAt time.Time
}

var (
	jwtSecret   = os.Getenv("JWT_SECRET")
	jwtDuration = time.Hour * 72
)

func NewJwtToken(info UserInfo) (JwtToken, error) {
	expireTime := time.Now().Add(jwtDuration)

	claims := &jwtClaims{
		info,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return *new(JwtToken), err
	}

	return JwtToken{
		Value:     signedToken,
		ExpiresAt: expireTime,
	}, nil
}
