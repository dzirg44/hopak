package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hopak/server/models"
	"github.com/hopak/server/utils"
)

// GetListOfPosts returns list of posts for a specific user
func (env *Env) GetListOfPosts(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value(userID).(uint)
	posts, err := models.GetPostsByUserID(env.DB, userID)

	if err != nil || len(posts) < 1 {
		utils.RespondError(w, http.StatusNotFound, utils.NoPostsForUser)
		return
	}

	utils.RespondJSON(w, http.StatusOK, &posts)
}

// CreatePost controller runs when user wants to create
func (env *Env) CreatePost(w http.ResponseWriter, req *http.Request) {

	var post models.Post
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&post)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		log.Print(err)
	}

	post.UserID = req.Context().Value(userID).(uint)
	post.CreatedAt = time.Now().Format("Mon Jan _2 15:04:05")

	err = models.CreateNewPost(env.DB, &post)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, &post)
}
