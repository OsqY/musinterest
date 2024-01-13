package authenticator

import (
	"context"
	"errors"
	"log"
	"oscar/musinterest/initializers"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func New() (*Authenticator, error) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Couldn't load env variables")
	}
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+config.Auth0Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     config.Auth0ClientID,
		ClientSecret: config.Auth0ClientSecret,
		RedirectURL:  config.Auth0CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
