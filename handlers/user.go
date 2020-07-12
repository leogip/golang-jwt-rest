package handlers

import (
	"net/http"

	"github.com/leogip/golang-jwt-rest/database"
	"github.com/leogip/golang-jwt-rest/lib/response"
	"github.com/leogip/golang-jwt-rest/models"
)

// GetAllUsers ...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var err error

	users, err = database.FindAllUser()
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Users not found", &users)
	}

	response.JSON(w, http.StatusOK, "You get all users", &users)
}