// Copyright 2020 CJ Harries
// Licensed under http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"

	"github.com/wizardsoftheweb/daily-programmer-cli/cmd"
)

func TestRootMain(t *testing.T) { TestingT(t) }

type MainSuite struct {
	errorMessage string
}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestMain(c *C) {
	var oldDailyProgCmd = &cobra.Command{}
	*oldDailyProgCmd = *cmd.DailyProgCmd
	dummy := func(cmd *cobra.Command, args []string) {}
	cmd.DailyProgCmd.SilenceErrors = true
	cmd.DailyProgCmd.DisableFlagParsing = true
	cmd.DailyProgCmd.PersistentPreRun = dummy
	cmd.DailyProgCmd.PreRun = dummy
	cmd.DailyProgCmd.Run = dummy
	cmd.DailyProgCmd.PostRun = dummy
	cmd.DailyProgCmd.PersistentPostRun = dummy
	c.Assert(
		func() {
			main()
		},
		Not(PanicMatches),
		"*",
	)
	*cmd.DailyProgCmd = *oldDailyProgCmd
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
