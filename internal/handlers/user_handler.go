package handlers

import (
	"log"
	"net/http"

	"github.com/Aleksandar-G/rss-aggregator/internal/models"
	"github.com/Aleksandar-G/rss-aggregator/internal/services"
	"github.com/Aleksandar-G/rss-aggregator/pkg"
	"github.com/go-chi/chi"
)

type UserHandler struct {
	userService *services.UserService
}

func (userHandler *UserHandler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on POST /v1/users/")
	// Create request body struct
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}

	body, err := pkg.DecodeRequestBody(r, params)
	if err != nil {
		pkg.RespondWithError(w, 401, err.Error())
	}

	dbUser, err := userHandler.userService.AddUserInDB(r.Context(), body.Name)
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database user object to models user
	user := models.DatabaseUserToUser(dbUser)

	// Return a successful response
	pkg.RespondWithJSON(w, http.StatusCreated, user)
}

func (userHandler *UserHandler) HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on DELETE /v1/users/{id}")
	userId := chi.URLParam(r, "id")
	if userId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	err := userHandler.userService.DeleteUserFromDB(r.Context(), userId)
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	pkg.RespondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "resource deleted",
	})
}

func (userHandler *UserHandler) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/users/{id}")
	userId := chi.URLParam(r, "id")
	if userId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	dbUser, err := userHandler.userService.FetchUserFromDB(r.Context(), userId)

	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database user to models user
	user := models.DatabaseUserToUser(dbUser)

	// Return response with user in body
	pkg.RespondWithJSON(w, http.StatusOK, user)
}

func (userHandler *UserHandler) HandlerListUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/users/")
	// Convert database user to models user
	var users []models.User

	dbUsers, err := userHandler.userService.ListUsersFromDB(r.Context())
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
	}
	for _, dbUser := range dbUsers {
		users = append(users, models.DatabaseUserToUser(dbUser))
	}

	// Return response with user in body
	pkg.RespondWithJSON(w, http.StatusOK, users)
}

// Return a new instance of UserHandler
func NewUserHandler() *UserHandler {
	userService := services.NewUserService()

	return &UserHandler{userService: userService}
}
