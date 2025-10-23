package handlers

import (
	"log"
	"net/http"

	"github.com/Aleksandar-G/rss-aggregator/internal/models"
	"github.com/Aleksandar-G/rss-aggregator/internal/services"
	"github.com/Aleksandar-G/rss-aggregator/pkg"
	"github.com/go-chi/chi"
)

type UserFeedHandler struct {
	userFeedService *services.UserFeedService
}

func (userFeedHandler *UserFeedHandler) HandlerCreateUserFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on POST /v1/user_feed/")
	// Create request body struct
	type parameters struct {
		UserId string `json:"user_id"`
		FeedId string `json:"feed_id"`
	}

	params := parameters{}

	body, err := pkg.DecodeRequestBody(r, params)
	if err != nil {
		pkg.RespondWithError(w, 401, err.Error())
	}

	dbUserFeed, err := userFeedHandler.userFeedService.AddUserFeedInDB(r.Context(), body.UserId, body.FeedId)
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database userFeed object to models feed
	userFeed := models.DatabaseUsersFeedToUserFeed(dbUserFeed)

	// Return a successful response
	pkg.RespondWithJSON(w, http.StatusCreated, userFeed)
}

func (userFeedHandler *UserFeedHandler) HandlerDeleteUserFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on DELETE /v1/user_feed/{id}")
	feedId := chi.URLParam(r, "id")
	if feedId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	err := userFeedHandler.userFeedService.DeleteUserFeedFromDB(r.Context(), feedId)
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

func (userFeedHandler *UserFeedHandler) HandlerGetUserFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/user_feed/{id}")
	feedId := chi.URLParam(r, "id")
	if feedId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	dbUserFeed, err := userFeedHandler.userFeedService.FetchUserFeedFromDB(r.Context(), feedId)

	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database user_feed to models userFeed
	userFeed := models.DatabaseUsersFeedToUserFeed(dbUserFeed)

	// Return response with userFeed in body
	pkg.RespondWithJSON(w, http.StatusOK, userFeed)
}

func (userFeedHandler *UserFeedHandler) HandlerListUserFeeds(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/user_feed/")
	// Convert database userFeed to models userFeed
	var userFeeds []models.UserFeed

	dbUserFeeds, err := userFeedHandler.userFeedService.ListUserFeedsFromDB(r.Context())
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
	}
	for _, dbUserFeed := range dbUserFeeds {
		userFeeds = append(userFeeds, models.DatabaseUsersFeedToUserFeed(dbUserFeed))
	}

	// Return response with feed in body
	pkg.RespondWithJSON(w, http.StatusOK, userFeeds)
}

// Return a new instance of UserFeedHandler
func NewUserFeedHandler() *UserFeedHandler {
	userFeedsService := services.NewUsersFeedsService()

	return &UserFeedHandler{userFeedService: userFeedsService}
}
