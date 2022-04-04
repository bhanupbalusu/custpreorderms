package user_auth

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	u "github.com/bhanupbalusu/custpreorderms/pkg/util/user_auth"
// 	form3jwt "github.com/form3tech-oss/jwt-go"
// 	//jwt "github.com/dgrijalva/jwt-go/v4"
// )

// var (
// 	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
// 	JwtSigningMethod = form3jwt.SigningMethodHS256.Name
// )

// func NewToken(userId string) (string, error) {
// 	claims := form3jwt.StandardClaims{
// 		Id:        userId,
// 		Issuer:    userId,
// 		IssuedAt:  time.Now().Unix(),
// 		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
// 	}
// 	token := form3jwt.NewWithClaims(form3jwt.SigningMethodHS256, claims)
// 	return token.SignedString(JwtSecretKey)
// }

// func validateSignedMethod(token *form3jwt.Token) (interface{}, error) {
// 	if _, ok := token.Method.(*form3jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}
// 	return JwtSecretKey, nil
// }

// func ParseToken(tokenString string) (*form3jwt.StandardClaims, error) {
// 	fmt.Println(tokenString)
// 	claims := new(form3jwt.StandardClaims)
// 	token, err := form3jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var ok bool
// 	claims, ok = token.Claims.(*form3jwt.StandardClaims)
// 	if !ok || !token.Valid {
// 		return nil, u.ErrInvalidAuthToken
// 	}
// 	return claims, nil
// }
