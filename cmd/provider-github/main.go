package main

import (
	githubProvider "github.com/ted-vo/provider-github/pkg/provider"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Provider: func() provider.Provider {
			return &githubProvider.GitHubRepository{}
		},
	})
}
