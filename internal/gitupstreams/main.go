package gitupstreams

import (
	"net/url"
	"strings"

	gurl "github.com/whilp/git-urls"
)

// GitURLHandler creates browser urls
type GitURLHandler struct {
	overwriteGitUpstream string
	handlers             []upstream
}

// NewGitURLHandler creates a new GitURLHandler with all known upstreams configured
func NewGitURLHandlerWithOverwrite(overwriteGitUpstream string) GitURLHandler {
	var h upstream
	switch overwriteGitUpstream {
	case "azure":
		h = AzureDevopsUpstream{}
	case "bitbucketorg":
		h = BitbucketOrgUpstream{}
	case "gitlab":
		h = GitlabUpstream{}
	default:
		h = GenericUpstream{}
	}
	return GitURLHandler{
		handlers: []upstream{h},
	}
}

// NewGitURLHandler creates a new GitURLHandler with all known upstreams configured
func NewGitURLHandler() GitURLHandler {
	return GitURLHandler{
		handlers: []upstream{
			AzureDevopsUpstream{},
			BitbucketOrgUpstream{},
			GitlabUpstream{},
		},
	}
}

type upstream interface {
	WillHandle(repoURL *url.URL) bool
	BranchURL(repoURL *url.URL, branch string) (string, error)
	PullRequestURL(repoURL *url.URL, branch string) (string, error)
	CIURL(repoURL *url.URL, branch string) (string, error)
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
	return GenericUpstream{}.BranchURL(url, branch)
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
	return GenericUpstream{}.PullRequestURL(url, branch)
}

// GetCIURL parses a git remote url, and create a url to be used in a browser
func (g GitURLHandler) GetCIURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL)
	if err != nil {
		return "", err
	}

	if domain != "" {
		url.Host = domain
	}

	for _, h := range g.handlers {
		if h.WillHandle(url) {
			return h.CIURL(url, branch)
		}
	}
	return GenericUpstream{}.CIURL(url, branch)
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
