package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/kevinburke/ssh_config"
	"github.com/mogensen/go-git-open/internal/gitupstreams"
	gurl "github.com/whilp/git-urls"

	"net/url"
)

func main() {
	gitRepo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	url, err := getURLFromGitRepo(gitRepo)
	if err != nil {
		log.Fatal(err)
	}

	openbrowser(url)
}

func getURLFromGitRepo(gitRepo *git.Repository) (string, error) {
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

		url, err := getBrowerURL(r.Config().URLs[0], domain, branch)
		if err != nil {
			return "", err
		}

		return url, nil
	}

	return "", fmt.Errorf("No remote url found")
}

func getBrowerURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL)
	if err != nil {
		return "", err
	}

	f, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	if err != nil {
		return "", err
	}
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return "", err
	}
	sshConfigDomain, _ := cfg.Get(url.Host, "HostName")

	if sshConfigDomain != "" && domain == "" {
		domain = sshConfigDomain
	}

	if strings.Contains(url.Host, "bitbucket.org") {
		url, err = gitupstreams.BitbucketOrgURL(url, branch)
		if err != nil {
			return "", err
		}
	} else if strings.Contains(url.Host, "azure.com") {
		url, err = gitupstreams.AzureURL(url, branch)
		if err != nil {
			return "", err
		}
	} else {
		url, err = gitupstreams.GenericURL(url, branch)
		if err != nil {
			return "", err
		}
	}

	fmt.Println("----")
	fmt.Printf("  - branch: %s\n", branch)
	fmt.Printf("  - domain: %s\n", domain)

	if domain != "" {
		url.Host = domain
	}
	fmt.Printf("%s\n", url)

	fmt.Println(remoteURL)
	return url.String(), nil
}

func getOverwriteDomain(gitRepo *git.Repository) string {

	conf, err := gitRepo.Config()
	if err != nil {
		panic(err)
	}

	sections := conf.Raw.Sections

	for _, s := range sections {
		if s.Name == "open" {
			return s.Options.Get("domain")
		}
	}

	return ""
}

func getURL(remote string) (*url.URL, error) {

	u, err := gurl.Parse(remote)
	if err != nil {
		return nil, err
	}

	browserURL := url.URL{
		Scheme: "https",
		Host:   u.Host,
		Path:   strings.TrimSuffix(u.Path, ".git"),
	}

	// If the URL is provided as "http", preserve that
	if u.Scheme == "http" {
		browserURL.Scheme = "http"
	}

	return &browserURL, nil
}

func openbrowser(url string) {
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
