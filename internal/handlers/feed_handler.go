package handlers

import (
	"log"
	"net/http"

	"github.com/Aleksandar-G/rss-aggregator/internal/models"
	"github.com/Aleksandar-G/rss-aggregator/internal/services"
	"github.com/Aleksandar-G/rss-aggregator/pkg"
	"github.com/go-chi/chi"
)

type FeedHandler struct {
	feedService *services.FeedService
}

func (feedHandler *FeedHandler) HandlerCreateFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on POST /v1/feeds/")
	// Create request body struct
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}

	body, err := pkg.DecodeRequestBody(r, params)
	if err != nil {
		pkg.RespondWithError(w, 401, err.Error())
	}

	dbFeed, err := feedHandler.feedService.AddFeedInDB(r.Context(), body.Name, body.URL)
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database feed object to models feed
	feed := models.DatabaseFeedToFeed(dbFeed)

	// Return a successful response
	pkg.RespondWithJSON(w, http.StatusCreated, feed)
}

func (feedHandler *FeedHandler) HandlerDeleteFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on DELETE /v1/feeds/{id}")
	feedId := chi.URLParam(r, "id")
	if feedId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	err := feedHandler.feedService.DeleteFeedFromDB(r.Context(), feedId)
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

func (feedHandler *FeedHandler) HandlerGetFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/feeds/{id}")
	feedId := chi.URLParam(r, "id")
	if feedId == "" {
		pkg.RespondWithError(w, http.StatusBadRequest, "Id is not passed")
		return
	}

	dbFeed, err := feedHandler.feedService.FetchFeedFromDB(r.Context(), feedId)

	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
		return
	}

	// Convert database feed to models feed
	feed := models.DatabaseFeedToFeed(dbFeed)

	// Return response with feed in body
	pkg.RespondWithJSON(w, http.StatusOK, feed)
}

func (feedHandler *FeedHandler) HandlerListFeeds(w http.ResponseWriter, r *http.Request) {
	log.Println("Request on GET /v1/feeds/")
	// Convert database feed to models feed
	var feeds []models.Feed

	dbFeeds, err := feedHandler.feedService.ListFeedsFromDB(r.Context())
	if err != nil {
		pkg.RespondWithError(w, 500, err.Error())
	}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, models.DatabaseFeedToFeed(dbFeed))
	}

	// Return response with feed in body
	pkg.RespondWithJSON(w, http.StatusOK, feeds)
}

// Return a new instance of FeedHandler
func NewFeedHandler() *FeedHandler {
	feedService := services.NewFeedService()

	return &FeedHandler{feedService: feedService}
}
