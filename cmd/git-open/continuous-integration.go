package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/mogensen/go-git-open/internal/gitupstreams"
	"github.com/spf13/cobra"
)

// ciCmd Short name for the CI command
var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Open Continuous Integration page (short name for continuous-integration)",
	Long:  `Open Continuous Integration page (short name for continuous-integration)`,
	Run:   continuousIntegrationCmd.Run,
}

// continuousIntegrationCmd reciesents the CI command
var continuousIntegrationCmd = &cobra.Command{
	Use:   "continuous-integration",
	Short: "Open Continuous Integration page",
	Long:  `Open Continuous Integration page`,
	Run: func(cmd *cobra.Command, args []string) {

		gitRepo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
		if err != nil {
			log.Fatal(err)
		}

		remote, domain, branch, _, err := getRepoInfo(gitRepo)
		if err != nil {
			log.Fatal(err)
		}

		guh := gitupstreams.NewGitURLHandler()
		desiredUpstream := getOverwriteGitUpstream(gitRepo)
		if desiredUpstream != "" {
			guh = gitupstreams.NewGitURLHandlerWithOverwrite(desiredUpstream)
		}

		url, err := guh.GetCIURL(remote, domain, branch)
		if err != nil {
			log.Fatal(err)
		}

		openBrowser(url)
	},
}
