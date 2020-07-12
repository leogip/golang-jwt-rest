package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// NewRefreshToken ...
func NewRefreshToken() string {
	r := make([]byte, 24)
	rand.Read(r)
	return base64.URLEncoding.EncodeToString(r)
}

// NewJWT ...
func NewJWT(uID string) (string, error) {
	authJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "uid": uID,
        "exp":  time.Now().Add(time.Second * 10).Unix(),
        "iat":  time.Now().Unix(),
    })

	return authJWT.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// UpdateJWT ...
func UpdateJWT(w http.ResponseWriter, uID string) (string, error) {
	accessToken, err := NewJWT(uID)
	if err != nil {
		logrus.WithError(err).Errorf("Can't create new JWT token")
		return "", err
	}

	w.Header().Set("Authorization", accessToken)
	return accessToken, nil
}

// ParseJWT ...
func ParseJWT(jwtToken string)  (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, err
}