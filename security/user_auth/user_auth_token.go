package user_auth

import (
	"fmt"
	"os"
	"time"

	u "github.com/bhanupbalusu/custpreorderms/pkg/util/user_auth"

	//jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	JwtUserAuthSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtUserAuthSigningMethod = jwt.SigningMethodHS256.Name
)

func GenerateNewToken(userId string) (string, error) {
	claims := jwt.RegisteredClaims{
		ID:        userId,
		Issuer:    userId,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	fmt.Println(token.SignedString(JwtUserAuthSecretKey))
	return token.SignedString(JwtUserAuthSecretKey)
}

func validateJWTSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return JwtUserAuthSecretKey, nil
}

func UserAuthParseToken(tokenString string) (*jwt.StandardClaims, error) {
	fmt.Println(tokenString)
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateJWTSignedMethod)
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, u.ErrInvalidAuthToken
	}
	return claims, nil
}
