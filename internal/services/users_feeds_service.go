package services

import (
	"context"
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/config"
	"github.com/Aleksandar-G/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type UserFeedService struct {
	dbConfig *config.APIConfig
}

// Add feed to database
func (userFeedService *UserFeedService) AddUserFeedInDB(ctx context.Context, user_id string, feed_id string) (database.UsersFeed, error) {
	// Create the feed
	dbFeed, err := userFeedService.dbConfig.DB.CreateUserFeed(ctx, database.CreateUserFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user_id,
		FeedID:    feed_id,
	})

	if err != nil {
		return database.UsersFeed{}, err
	}

	return dbFeed, nil
}

func (userFeedService *UserFeedService) DeleteUserFeedFromDB(ctx context.Context, id string) error {

	err := userFeedService.dbConfig.DB.DeleteUserFeed(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (userFeedService *UserFeedService) FetchUserFeedFromDB(ctx context.Context, id string) (database.UsersFeed, error) {
	dbFeed, err := userFeedService.dbConfig.DB.GetUserFeedById(ctx, id)
	if err != nil {
		return database.UsersFeed{}, err
	}
	return dbFeed, nil
}

func (userFeedService *UserFeedService) ListUserFeedsFromDB(ctx context.Context) ([]database.UsersFeed, error) {
	// Fetch the feed from the database
	dbFeeds, err := userFeedService.dbConfig.DB.ListUserFeeds(ctx)
	if err != nil {
		return nil, err
	}
	return dbFeeds, nil
}

// Return a new instance on FeedService
func NewUsersFeedsService() *UserFeedService {
	apiCfg := config.NewAPIConfig()

	return &UserFeedService{
		dbConfig: apiCfg,
	}
}
