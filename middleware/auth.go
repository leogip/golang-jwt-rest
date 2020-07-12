package middleware

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/leogip/golang-jwt-rest/database"
	"github.com/leogip/golang-jwt-rest/lib/response"
	"github.com/leogip/golang-jwt-rest/logger"
	"github.com/leogip/golang-jwt-rest/security"
)

var log = logger.New("Auth")

// AuthRequired ...
func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			response.JSON(w, http.StatusUnauthorized, "Header Authorization empty, JWT not found", false)
			return
		}

		jwtToken := authHeader[1]
		accessToken, err := security.ParseJWT(jwtToken)

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
					claims, _ := accessToken.Claims.(jwt.MapClaims)
					uID, _ := primitive.ObjectIDFromHex(claims["uid"].(string))
					refresh, err := database.FindTokenByUser(uID)
					if err != nil {
						response.JSON(w, http.StatusUnauthorized, "Refresh JWT not found", false)
						return
					}
					if refresh.ExpriedAt.UnixNano() <= time.Now().UnixNano() {
						response.JSON(w, http.StatusUnauthorized, "Refresh JWT is expried", false)
						return
					}
					_, err = security.UpdateJWT(w, claims["uid"].(string))
						if err != nil {
							log.WithError(err).Errorf("Can't update JWT")
							response.JSON(w, http.StatusServiceUnavailable, "Can't update JWT", false)
							return
						}
						w.Header().Set("userId", claims["uid"].(string))
			
						next.ServeHTTP(w, r)
				} else {
					response.JSON(w, http.StatusUnauthorized, "JWT Invalid", false)
					return
				}
			}
		}

		if _, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
			claims, _ := accessToken.Claims.(jwt.MapClaims)
			w.Header().Set("userId", claims["uid"].(string))
			next.ServeHTTP(w, r)
		}
	})
}

