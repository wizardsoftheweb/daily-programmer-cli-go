package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// The is the git-wiz version only. Other components may have
// different versions.
var PackageVersion = "undefined"

// This is the flag to check the command version
var ShowVersion bool

// The verbosity flag is a count flag, ie the more there are the more verbose
// it gets.
var VerbosityFlagValue int

// This flag is populated with a GitHub personal access token with at least the
// public_repo scope
var GithubPersonalAccessToken string

func init() {
	DailyProgCmd.PersistentFlags().BoolVarP(
		&ShowVersion,
		"version",
		"V",
		false,
		"Prints the app version",
	)
	DailyProgCmd.PersistentFlags().CountVarP(
		&VerbosityFlagValue,
		"verbose",
		"v",
		"Increases application verbosity",
	)
	DailyProgCmd.PersistentFlags().StringVar(
		&GithubPersonalAccessToken,
		"githubpat",
		os.Getenv("WOTW_DAILY_PROG_GITHUB_TOKEN"),
		"GH PAT with at least public_repo scope",
	)
	_ = DailyProgCmd.PersistentFlags().MarkHidden("githubpat")
}

// This is the primary cmd runner and exposes git-wiz
func Execute() error {
	return DailyProgCmd.Execute()
}

// daily-prog has no base functionality. It must be used with subcommands.
var DailyProgCmd = &cobra.Command{
	Use:   "daily-prog",
	Short: "i have no idea what im doing",
	Long:  "opinionated r/DailyProgrammer tooling",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ShowVersion {
			cmd.Printf("%s version %s\n", cmd.Use, PackageVersion)
			cmd.TraverseChildren = false
			cmd.Run = func(cmd *cobra.Command, args []string) {}
		}
	},
	Run: HelpOnly,
}

// This is a catch-all error handler that kills the program when an
// error occurs.
func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
