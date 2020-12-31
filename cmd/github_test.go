package cmd

import (
	"github.com/google/go-github/github"
	"net/http"
	"testing"

	. "gopkg.in/check.v1"
)

func TestGithub(t *testing.T) { TestingT(t) }

const sharedPat = "qqq123"
var oldNewClient func(httpClient *http.Client) *github.Client
var RepositoriesServiceType RepositoriesService

type GithubSuite struct {
	BaseSuite
}

func mockNewClient(httpClient *http.Client) *github.Client {
	return &github.Client{}
}

var _ = Suite(&GithubSuite{})

func (s *GithubSuite) SetUpTest(c *C) {
	oldNewClient = zGithubNewClient
	zGithubNewClient = mockNewClient
}

func (s *GithubSuite) TearDownTest(c *C) {
	zGithubNewClient = oldNewClient
}

func (s *GithubSuite) TestGetRepositoriesService(c *C) {
	repoSvc := getRepositoriesService(sharedPat)
	c.Assert(repoSvc, Implements, &RepositoriesServiceType)
}

func (s *GithubSuite) TestDoesRepoExist(c *C) {

}

func (s *GithubSuite) TestEnsureRepoExists(c *C) {

}
