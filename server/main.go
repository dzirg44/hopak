package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"github.com/hopak/server/config"
	"github.com/hopak/server/controllers"
	"github.com/hopak/server/driver"
)

func main() {

	// Init DB connection
	config := config.GetConfig()
	db := driver.ConnectSQL(config)
	defer db.Close()

	controllersEnv := &controllers.Env{DB: db}
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.AllowAll().Handler)

	router.Post("/api/auth/signup", controllersEnv.SignUp)
	router.Post("/api/auth/signin", controllersEnv.SignIn)

	router.Group(func(router chi.Router) {
		router.Use(controllers.IsAuthorized)
		router.Use(controllers.EnsureCorrectUser)
		router.Get("/api/users/{user_id}", controllersEnv.GetUser)
		router.Get("/api/users/{user_id}/posts", controllersEnv.GetListOfPosts)
		router.Post("/api/users/{user_id}/posts", controllersEnv.CreatePost)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
