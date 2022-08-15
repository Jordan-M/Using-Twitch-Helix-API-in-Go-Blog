package main

import (
	"fmt"

	"github.com/nicklaw5/helix/v2"
)

var twitchClient *helix.Client
var err error

func InitTwitchClient() {
	twitchClient, err = helix.NewClient(&helix.Options{
		ClientID:     "<Replace Me>",
		ClientSecret: "<Replace Me>",
	})

	if err != nil {
		panic(err)
	}

	accessTokenResponse, accessTokenResponseError := twitchClient.RequestAppAccessToken([]string{})
	if accessTokenResponseError != nil || accessTokenResponse.StatusCode != 200 {
		panic(accessTokenResponseError)
	}

	twitchClient.SetAppAccessToken(accessTokenResponse.Data.AccessToken)
}

func main() {
	InitTwitchClient()

	userResponse, userResponseError := twitchClient.GetUsers(&helix.UsersParams{
		Logins: []string{"xqc", "sodapoppin"},
	})

	if userResponseError != nil {
		fmt.Printf("Error getting users: %s", userResponseError)
	}

	for _, user := range userResponse.Data.Users {
		fmt.Printf("DisplayName: %s Id: %s\n", user.DisplayName, user.ID)
	}
}
