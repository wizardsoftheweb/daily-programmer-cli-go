package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/github"

	. "gopkg.in/check.v1"
)

func TestGithub(t *testing.T) { TestingT(t) }

const sharedPat = "qqq123"

var panicError = errors.New("whoops")

var oldNewClient func(httpClient *http.Client) *github.Client
var RepositoriesServiceType RepositoriesService

type GithubSuite struct {
	BaseSuite
	GetRepoSvc GetRepoSvcMock
}

func mockNewClient(httpClient *http.Client) *github.Client {
	return &github.Client{}
}

type GetRepoSvcMock struct{}

const repoExistsName string = "repo-exists"
const repoDneName string = "repo-dne"

func (rs GetRepoSvcMock) Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	if repoExistsName == repo {
		return nil, nil, nil
	}
	if repoDneName == repo {
		return nil, nil, errors.New(fmt.Sprintf("GET https://api.github.com/repos/qqq/%s: 404 Not Found []", repoDneName))
	}
	return nil, nil, panicError
}

func (rs GetRepoSvcMock) Create(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error) {
	return nil, nil, nil
}

var _ = Suite(&GithubSuite{})

func (s *GithubSuite) SetUpTest(c *C) {
	oldNewClient = zGithubNewClient
	zGithubNewClient = mockNewClient
	s.GetRepoSvc = GetRepoSvcMock{}
}

func (s *GithubSuite) TearDownTest(c *C) {
	zGithubNewClient = oldNewClient
}

func (s *GithubSuite) TestGetRepositoriesService(c *C) {
	repoSvc := getRepositoriesService(sharedPat)
	c.Assert(repoSvc, Implements, &RepositoriesServiceType)
}

func (s *GithubSuite) TestDoesRepoExist(c *C) {
	c.Assert(doesRepoExist(s.GetRepoSvc, repoExistsName), Equals, true)
	c.Assert(doesRepoExist(s.GetRepoSvc, repoDneName), Equals, false)
	c.Assert(
		func() { doesRepoExist(s.GetRepoSvc, "qqq") },
		Panics,
		panicError,
	)

}

func (s *GithubSuite) TestEnsureRepoExists(c *C) {

}
