package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lntvan166/e2tech-auth-svc/internal/db"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Phone string
	Role  string
}

func (w *JwtWrapper) GenerateToken(user db.User) (signedToken string, err error) {
	expiredAt := time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix()
	if user.Role == ADMIN {
		expiredAt = time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours*24*7)).Unix()
	}
	claims := &jwtClaims{
		Id:    user.ID,
		Phone: user.Phone,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

}
