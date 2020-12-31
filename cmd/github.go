package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Call out the github.RepositoriesService functionality we need so we can mock
// mock it during testing
type RepositoriesService interface {
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
	Create(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error)
}

// Expose github.NewClient functionality to tests so we can mock it
var zGithubNewClient = github.NewClient

// Creates a new GH client using a passed-in token
// See https://github.com/google/go-github#authentication
func getRepositoriesService(personalAccessToken string) *github.RepositoriesService {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return zGithubNewClient(tc).Repositories
}

// Checks to see if the given `repoName` exists
func doesRepoExist(repoSvc RepositoriesService, repoName string) bool {
	ctx := context.Background()
	_, _, err := repoSvc.Get(ctx, "", repoName)
	if nil != err {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false
		}
		whereErrorsGoToDie(err)
	}
	return true
}

// Ensures the given `repoName` exists; creates if DNE and does nothing otherwise
func ensureRepoExists(repoSvc RepositoriesService, repoName string) {
	if !doesRepoExist(repoSvc, repoName) {
		ctx := context.Background()
		_, _, err := repoSvc.Create(
			ctx,
			"",
			&github.Repository{
				Name: github.String(repoName),
			},
		)
		whereErrorsGoToDie(err)
		fmt.Print(err)
	}
}
