package client

import (
	"errors"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/srt32/prpluck/config"
)

func NewClient(userConfig config.Config) (*github.Client, error) {
	userToken := userConfig.Token()

	if userToken == "" {
		return &github.Client{}, errors.New("GitHub token is required")
	}

	t := &oauth.Transport{Token: &oauth.Token{AccessToken: userToken}}

	return github.NewClient(t.Client()), nil
}
