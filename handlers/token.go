package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/leogip/golang-jwt-rest/database"
	"github.com/leogip/golang-jwt-rest/lib/response"
	"github.com/leogip/golang-jwt-rest/logger"
	"github.com/leogip/golang-jwt-rest/models"
	"github.com/leogip/golang-jwt-rest/security"
)

var log = logger.New("Tokens")

// GetAllTokens ...
func GetAllTokens(w http.ResponseWriter, r *http.Request) {
	var tokens []models.Token
	var err error

	tokens, err = database.FindAllTokens()
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Users not found", &tokens)
	}

	response.JSON(w, http.StatusOK, "You get all users", &tokens)
}

// GetTokens ...
func GetTokens(w http.ResponseWriter, r *http.Request) {
	var token models.Token
	params := mux.Vars(r)
	uID := params["userId"]
	userID, _ := primitive.ObjectIDFromHex(uID)

	user, err := database.FindUserByID(userID)
	if err != nil {
		log.Warn(err)
		response.JSON(w, http.StatusBadRequest, "Invalid user ID", false)
		return
	}
	token, err = database.FindTokenByUser(user.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			token.ID = primitive.NewObjectID()
			token.UserID = user.ID
			token.Token = security.NewRefreshToken()
			token.ExpriedAt = time.Now().UTC()
		
			if err = database.InsertToken(token); err != nil {
				log.Warn(err)
				response.JSON(w, http.StatusServiceUnavailable, "Server Error", false)
				return
			}
		}
	}

	var accessToken string
	accessToken, err = security.UpdateJWT(w,uID)
	if err != nil {
		log.Warn(err)
		response.JSON(w, http.StatusNoContent, "Can't update JWT", false)
		return
	}

	response.JSON(w, http.StatusOK, "Created Refresh Token, in database. Access Token saved on localstorage", &accessToken)
}
// RefreshTokens ...
func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	uID := w.Header().Get("userId")
	userID, _ := primitive.ObjectIDFromHex(uID)

	token, err := database.FindTokenByUser(userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			token.ID = primitive.NewObjectID()
			token.UserID = userID
			token.Token = security.NewRefreshToken()
			
			if err = database.InsertToken(token); err != nil {
				response.JSON(w, http.StatusServiceUnavailable, "Server Error", false)
				return
			}
		}
	} else {
		err = database.UpdateToken(token.ID, security.NewRefreshToken())
		if err != nil {
			response.JSON(w, http.StatusServiceUnavailable, "Server Error: err update token", false)
		}
	}

	var accessToken string
	accessToken, err = security.UpdateJWT(w, uID)
	if err != nil {
		log.WithError(err).Errorf("Can't update JWT")
	}

	response.JSON(w, http.StatusCreated, "Refresh and Access Tokens updated", accessToken)
}
// RemoveToken ...
func RemoveToken(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tkID := params["tokenId"]
	tokenID, _ := primitive.ObjectIDFromHex(tkID)

	err := database.RemoveTokenByID(tokenID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response.JSON(w, http.StatusBadRequest, "Invalid token ID", false)
			return
		}
		response.JSON(w, http.StatusServiceUnavailable, "Server Error", false)
		return
	}

	response.JSON(w, http.StatusOK, "Token by id:" + tkID + " remove", true)
}
// RemoveAllTokens ...
func RemoveAllTokens(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uID := params["userId"]
	userID, _ := primitive.ObjectIDFromHex(uID)

	err := database.RemoveTokenByUserID(userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response.JSON(w, http.StatusBadRequest, "Not found tokens with this user", false)
			return
		}
		response.JSON(w, http.StatusServiceUnavailable, "Server Error", false)
		return
	}

	response.JSON(w, http.StatusOK, "All Tokens by userId:" + uID + " remove", true)
}
