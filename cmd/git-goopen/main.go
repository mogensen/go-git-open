package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/go-git/go-git/v5"
	"github.com/kevinburke/ssh_config"
	"github.com/mogensen/go-git-open/internal/gitupstreams"
	gurl "github.com/whilp/git-urls"
)

func main() {
	gitRepo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		log.Fatal(err)
	}

	url, err := getURLFromGitRepo(gitRepo)
	if err != nil {
		log.Fatal(err)
	}

	openBrowser(url)
}

func getURLFromGitRepo(gitRepo *git.Repository) (string, error) {

	guh := gitupstreams.NewGitURLHandler()

	list, err := gitRepo.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range list {

		// if domain is set in git options we override with this
		domain := getOverwriteDomain(gitRepo)
		branch := ""

		h, err := gitRepo.Head()
		if err != nil {
			return "", err
		}
		if h.Name().IsBranch() {
			branch = h.Name().Short()
		}

		url, err := guh.GetBrowerURL(r.Config().URLs[0], domain, branch)
		if err != nil {
			return "", err
		}

		return url, nil
	}

	return "", fmt.Errorf("No remote url found")
}

func getOverwriteDomain(gitRepo *git.Repository) string {

	// If we cannot find any config, just give up
	conf, err := gitRepo.Config()
	if err != nil {
		return ""
	}

	// If we find a open.domain config we use this
	for _, s := range conf.Raw.Sections {
		if s.Name == "open" {
			return s.Options.Get("domain")
		}
	}

	// Lookup if the domain is a ssl alias
	list, err := gitRepo.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	url, err := gurl.Parse(list[0].Config().URLs[0])
	if err != nil {
		return ""
	}

	// Lookup Hostname alias in ssh config, empty if none is found
	sshConf := ssh_config.DefaultUserSettings
	sshConf.IgnoreErrors = true
	return sshConf.Get(url.Host, "HostName")
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
