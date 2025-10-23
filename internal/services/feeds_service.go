package services

import (
	"context"
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/config"
	"github.com/Aleksandar-G/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedService struct {
	dbConfig *config.APIConfig
}

// Add feed to database
func (feedService *FeedService) AddFeedInDB(ctx context.Context, name string, url string) (database.Feed, error) {
	// Create the feed
	dbFeed, err := feedService.dbConfig.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
	})

	if err != nil {
		return database.Feed{}, err
	}

	return dbFeed, nil
}

func (feedService *FeedService) DeleteFeedFromDB(ctx context.Context, feedId string) error {

	err := feedService.dbConfig.DB.DeleteUser(ctx, feedId)
	if err != nil {
		return err
	}
	return nil
}

func (feedService *FeedService) FetchFeedFromDB(ctx context.Context, feedId string) (database.Feed, error) {
	dbFeed, err := feedService.dbConfig.DB.GetFeedById(ctx, feedId)
	if err != nil {
		return database.Feed{}, err
	}
	return dbFeed, nil
}

func (feedService *FeedService) ListFeedsFromDB(ctx context.Context) ([]database.Feed, error) {
	// Fetch the feed from the database
	dbFeeds, err := feedService.dbConfig.DB.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}
	return dbFeeds, nil
}

// Return a new instance on FeedService
func NewFeedService() *FeedService {
	apiCfg := config.NewAPIConfig()

	return &FeedService{
		dbConfig: apiCfg,
	}
}
