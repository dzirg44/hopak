package controllers

import (
	"fmt"
	"net/http"

	"github.com/hopak/server/models"
	"github.com/hopak/server/utils"
	"github.com/jinzhu/gorm"
)

// Env environment structure for our api controllers
type Env struct {
	DB *gorm.DB
}

// GetUser returns the user
func (env *Env) GetUser(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value(userID).(uint)
	user := models.User{UserID: userID}
	userNotFound := models.UserByID(env.DB, &user)

	if userNotFound != false {
		utils.RespondError(w, http.StatusNotFound, utils.UserNotFound)
		return
	}

	posts, err := models.GetPostsByUserID(env.DB, user.UserID)

	if err != nil {
		fmt.Printf("No posts found for the user: %s", user.Username)
	}
	user.Posts = posts
	utils.RespondJSON(w, http.StatusOK, &user)
}
