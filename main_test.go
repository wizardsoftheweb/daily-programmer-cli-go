package main

import (
	"errors"
	"testing"

	. "gopkg.in/check.v1"
)

func TestRootMain(t *testing.T) { TestingT(t) }

type MainSuite struct {
	errorMessage string
}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestMain(c *C) {
	c.Assert(
		func() {
			main()
		},
		Not(PanicMatches),
		"*",
	)
}

func (s *MainSuite) TestWhereErrorsGoToDie(c *C) {
	s.errorMessage = "qqq"
	c.Assert(
		func() {
			whereErrorsGoToDie(errors.New(s.errorMessage))
		},
		PanicMatches,
		s.errorMessage,
	)
}

