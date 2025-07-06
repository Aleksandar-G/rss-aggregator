package services

import (
	"context"
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/config"
	"github.com/Aleksandar-G/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type UserService struct {
	dbConfig *config.APIConfig
}

// Add user to database
func (userService *UserService) AddUserInDB(ctx context.Context, name string) (database.User, error) {
	// Create the user
	dbUser, err := userService.dbConfig.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})

	if err != nil {
		return database.User{}, err
	}

	return dbUser, nil
}

func (userService *UserService) DeleteUserFromDB(ctx context.Context, userId string) error {

	err := userService.dbConfig.DB.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) FetchUserFromDB(ctx context.Context, userId string) (database.User, error) {
	dbUser, err := userService.dbConfig.DB.GetUserById(ctx, userId)
	if err != nil {
		return database.User{}, err
	}
	return dbUser, nil
}

func (userService *UserService) ListUsersFromDB(ctx context.Context) ([]database.User, error) {
	// Fetch the user from the database
	dbUsers, err := userService.dbConfig.DB.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return dbUsers, nil
}

// Return a new instance on UserService
func NewUserService() *UserService {
	apiCfg := config.NewAPIConfig()

	return &UserService{
		dbConfig: apiCfg,
	}
}
