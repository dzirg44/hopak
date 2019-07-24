package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"

	"github.com/dzirg44/hopak/server/models"
	"github.com/dzirg44/hopak/server/utils"
	"golang.org/x/crypto/bcrypt"
)

// key is context key structure
type key int

// keep userID saved
const userID key = 0

// SignUp registers a new user
func (env *Env) SignUp(w http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// hash and salt password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPass)

	if len(user.Username) < 1 || len(user.Password) < 1 || len(user.Email) < 1 {
		utils.RespondError(w, http.StatusBadRequest, "missing required body parameters")
		return
	}

	// Check if user exist
	// if user not found return proper message
	userNotFound := models.IsUserExist(env.DB, user.Email)

	if userNotFound == false {
		utils.RespondError(w, http.StatusConflict, utils.UserAlreadyRegistered)
		return
	}

	// Add to a database
	env.DB.Create(&user)

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	authResponse := models.AuthResponse{UserID: user.UserID, Username: user.Username, Token: token}

	// Return the required data and token
	utils.RespondJSON(w, http.StatusOK, &authResponse)

}

// SignIn logins user successfully if password and username is correct
func (env *Env) SignIn(w http.ResponseWriter, req *http.Request) {
	var user models.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	candidatePass := user.Password

	// Check if user exist
	// if user found, continue found return proper message
	// if user not found, stop and return the proper message
	userNotFound := models.IsUserExist(env.DB, user.Email)

	if userNotFound == false {
		env.DB.Where("email = ?", user.Email).First(&user)
	} else {
		utils.RespondError(w, http.StatusNotFound, utils.UserNotFound)
		return
	}

	// Compare passed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidatePass))

	// If failed, return the proper message
	if err != nil {
		log.Println(err)
		utils.RespondError(w, http.StatusUnauthorized, utils.InvalidPassword)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	authResponse := models.AuthResponse{UserID: user.UserID, Username: user.Username, Token: token}

	// Return the required data and token
	utils.RespondJSON(w, http.StatusOK, &authResponse)
}

// IsAuthorized checks if the user is Authorized
func IsAuthorized(next http.Handler) http.Handler {
	// Create a new Middleware
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenHeader := req.Header.Get("Authorization") // Grab the token from the header

		// If token is missing, returns Unauthorized message
		if tokenHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, utils.MissingToken)
			return
		}

		// The token format is `Bearer {token}`, should match this format
		splittedToken := strings.Split(tokenHeader, " ")
		if len(splittedToken) != 2 {
			utils.RespondError(w, http.StatusForbidden, utils.InvalidToken)
			return
		}

		// Get the token part
		tokenPart := splittedToken[1]
		tokenObj := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tokenObj, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.MySecretKey), nil
		})

		if err != nil {
			utils.RespondError(w, http.StatusForbidden, "Malformed authentication token")
			return
		}

		if !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, utils.Unauthorized)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(req.Context(), userID, tokenObj.UserID)
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req) //proceed in the middleware chain!
	})
}

// EnsureCorrectUser ensures that user_id in token is the same with request param id
func EnsureCorrectUser(next http.Handler) http.Handler {
	// Create a new Middleware
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		paramID, err := strconv.ParseUint(chi.URLParam(req, "user_id"), 10, 32)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctxUserID := req.Context().Value(userID).(uint)

		if uint(paramID) != ctxUserID {
			utils.RespondError(w, http.StatusUnauthorized, utils.Unauthorized)
			return
		}

		next.ServeHTTP(w, req) //proceed in the middleware chain!
	})
}
