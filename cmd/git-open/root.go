package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/mogensen/go-git-open/internal/gitupstreams"
	"github.com/spf13/cobra"
)

// Print overrides the browser, so the resulting url is printed to stdout
var Print bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "git-open",
	Long: `This is an extension for the git-cli, that allows you to open any git repository in your browser.`,
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

		url, err := guh.GetBrowerURL(remote, domain, branch)
		if err != nil {
			log.Fatal(err)
		}

		openBrowser(url)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Print, "print", "p", false, "print the browser url, instead of opening it")
	rootCmd.AddCommand(prCmd)
	rootCmd.AddCommand(pullRequestCmd)
	rootCmd.AddCommand(ciCmd)
	rootCmd.AddCommand(continuousIntegrationCmd)
}
