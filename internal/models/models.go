package models

import (
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/database"
)

type User struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

type Feed struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
	URL       string      `json:"url"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
	}
}

type UserFeed struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User_id   interface{} `json:"user_id"`
	Feed_id   interface{} `json:"feed_id"`
}

func DatabaseUsersFeedToUserFeed(dbUsers_feeds database.UsersFeed) UserFeed {
	return UserFeed{
		ID:        dbUsers_feeds.ID,
		CreatedAt: dbUsers_feeds.CreatedAt,
		UpdatedAt: dbUsers_feeds.UpdatedAt,
		User_id:   dbUsers_feeds.UserID,
		Feed_id:   dbUsers_feeds.FeedID,
	}
}
