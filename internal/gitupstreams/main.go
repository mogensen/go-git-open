package gitupstreams

import (
	"fmt"
	"net/url"
	"strings"

	gurl "github.com/whilp/git-urls"
)

// GitURLHandler creates browser urls
type GitURLHandler struct {
	handlers []upstream
}

// NewGitURLHandler creates a new GitURLHandler with all known upstreams configured
func NewGitURLHandler() GitURLHandler {
	return GitURLHandler{
		handlers: []upstream{
			AzureDevopsUpstream{},
			BitbucketOrgUpstream{},
			GenericUpstream{},
		},
	}
}

type upstream interface {
	WillHandle(repoURL *url.URL) bool
	BranchURL(repoURL *url.URL, branch string) (string, error)
	PullRequestURL(repoURL *url.URL, branch string) (string, error)
}

// GetBrowerURL parses a git remote url, and create a url to be used in a browser
func (g GitURLHandler) GetBrowerURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL)
	if err != nil {
		return "", err
	}

	if domain != "" {
		url.Host = domain
	}

	for _, h := range g.handlers {
		if h.WillHandle(url) {
			return h.BranchURL(url, branch)
		}
	}
	// This should never happen, as the generic handler will try to handle anything
	return "", fmt.Errorf("Found no handlers for url: %s", url.String())
}

// GetPullRequestURL parses a git remote url, and create a url to be used in a browser
func (g GitURLHandler) GetPullRequestURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL)
	if err != nil {
		return "", err
	}

	if domain != "" {
		url.Host = domain
	}

	for _, h := range g.handlers {
		if h.WillHandle(url) {
			return h.PullRequestURL(url, branch)
		}
	}
	// This should never happen, as the generic handler will try to handle anything
	return "", fmt.Errorf("Found no handlers for url: %s", url.String())
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
