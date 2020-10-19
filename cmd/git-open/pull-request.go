package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/mogensen/go-git-open/internal/gitupstreams"
	"github.com/spf13/cobra"
)

// prCmd Short name for the pullRequest command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Open Pull Request page (short name for pull-request)",
	Long:  `Open Pull Request page (short name for pull-request)`,
	Run:   pullRequestCmd.Run,
}

// pullRequestCmd represents the pullRequest command
var pullRequestCmd = &cobra.Command{
	Use:   "pull-request",
	Short: "Open Pull Request page",
	Long:  `Open Pull Request page`,
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
		url, err := guh.GetPullRequestURL(remote, domain, branch)
		if err != nil {
			log.Fatal(err)
		}

		openBrowser(url)
	},
}
