package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
)

func GetConnection() (*auth.Client, error) {
	opt := option.WithCredentialsFile("internal/infrastructure/auth/firebase/credentials.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("error initializing firebase connection:", err)
		return nil, err
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("error initializing firebase connection:", err)
		return nil, err
	}

	return authClient, nil
}
