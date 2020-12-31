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
	RepoSvc RepoSvcMock
}

func mockNewClient(httpClient *http.Client) *github.Client {
	return &github.Client{}
}

type RepoSvcMock struct{}

const repoExistsName string = "repo-exists"
const repoDneName string = "repo-dne"

var callCount int

func (rs RepoSvcMock) Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	if repoExistsName == repo {
		return nil, nil, nil
	}
	if repoDneName == repo {
		return nil, nil, errors.New(fmt.Sprintf("GET https://api.github.com/repos/qqq/%s: 404 Not Found []", repoDneName))
	}
	return nil, nil, panicError
}

func (rs RepoSvcMock) Create(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error) {
	callCount += 1
	return nil, nil, nil
}

var _ = Suite(&GithubSuite{})

func (s *GithubSuite) SetUpTest(c *C) {
	oldNewClient = zGithubNewClient
	zGithubNewClient = mockNewClient
	s.RepoSvc = RepoSvcMock{}
	callCount = 0
}

func (s *GithubSuite) TearDownTest(c *C) {
	zGithubNewClient = oldNewClient
}

func (s *GithubSuite) TestGetRepositoriesService(c *C) {
	repoSvc := getRepositoriesService(sharedPat)
	c.Assert(repoSvc, Implements, &RepositoriesServiceType)
}

func (s *GithubSuite) TestDoesRepoExist(c *C) {
	c.Assert(doesRepoExist(s.RepoSvc, repoExistsName), Equals, true)
	c.Assert(doesRepoExist(s.RepoSvc, repoDneName), Equals, false)
	c.Assert(
		func() { doesRepoExist(s.RepoSvc, "qqq") },
		Panics,
		panicError,
	)

}

func (s *GithubSuite) TestEnsureRepoExists(c *C) {
	ensureRepoExists(s.RepoSvc, repoExistsName)
	c.Assert(callCount, Equals, 0)
	ensureRepoExists(s.RepoSvc, repoDneName)
	c.Assert(callCount, Equals, 1)
}
