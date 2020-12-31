package cmd

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	. "gopkg.in/check.v1"
)

func TestDailyProg(t *testing.T) { TestingT(t) }

type DailyProgSuite struct {
	BaseSuite
}

var _ = Suite(&DailyProgSuite{})

func (s *DailyProgSuite) TestExecute(c *C) {
	var oldDailyProgCmd = &cobra.Command{}
	*oldDailyProgCmd = *DailyProgCmd
	dummy := func(cmd *cobra.Command, args []string) {}
	DailyProgCmd.SilenceErrors = true
	DailyProgCmd.DisableFlagParsing = true
	DailyProgCmd.PersistentPreRun = dummy
	DailyProgCmd.PreRun = dummy
	DailyProgCmd.Run = dummy
	DailyProgCmd.PostRun = dummy
	DailyProgCmd.PersistentPostRun = dummy
	err := Execute()
	c.Assert(err, IsNil)
	*DailyProgCmd = *oldDailyProgCmd
}

func (s *DailyProgSuite) TestWhereErrorsGoToDie(c *C) {
	s.errorMessage = "qqq"
	c.Assert(
		func() {
			whereErrorsGoToDie(errors.New(s.errorMessage))
		},
		PanicMatches,
		s.errorMessage,
	)
}
