package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty paramters")
	}
	// does user exist for the Username
	userExist, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("there was an error checking if the user exist %w", err)
	}

	if userExist {
		return fmt.Errorf("user with this username exist")
	}

	// insert the user since it doesn't exist
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("there was an error inserting the user %w", err)
	}

	return nil
}
