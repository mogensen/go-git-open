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

// NewGitURLHandlerWithOverwrite creates a new GitURLHandler containing only the specified upstream
func NewGitURLHandlerWithOverwrite(overwriteGitUpstream string) GitURLHandler {
	var h upstream

	switch overwriteGitUpstream {
	case "azure":
		h = AzureDevopsUpstream{}
	case "bitbucketorg":
		h = BitbucketOrgUpstream{}
	case "bitbucket":
		h = BitbucketUpstream{}
	case "gitlab":
		h = GitlabUpstream{}
	default:
		h = GenericUpstream{}
	}
	return GitURLHandler{
		handlers:             []upstream{h},
		overwriteGitUpstream: overwriteGitUpstream,
	}
}

// NewGitURLHandler creates a new GitURLHandler with all known upstreams configured
func NewGitURLHandler() GitURLHandler {
	return GitURLHandler{
		handlers: []upstream{
			AzureDevopsUpstream{},
			BitbucketOrgUpstream{},
			BitbucketUpstream{},
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
	url, err := getURL(remoteURL, domain)
	if err != nil {
		return "", err
	}

	return g.getProvider(url).BranchURL(url, branch)

}

// GetPullRequestURL parses a git remote url, and create a url to be used in a browser
func (g GitURLHandler) GetPullRequestURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL, domain)
	if err != nil {
		return "", err
	}

	return g.getProvider(url).PullRequestURL(url, branch)
}

// GetCIURL parses a git remote url, and create a url to be used in a browser
func (g GitURLHandler) GetCIURL(remoteURL string, domain, branch string) (string, error) {
	url, err := getURL(remoteURL, domain)
	if err != nil {
		return "", err
	}

	return g.getProvider(url).CIURL(url, branch)
}

func (g GitURLHandler) getProvider(url *url.URL) upstream {
	// If we have an overwrite git upstream, we dont ask if the provider will handle the url
	if g.overwriteGitUpstream != "" {
		return g.handlers[0]
	}

	for _, h := range g.handlers {
		if h.WillHandle(url) {
			return h
		}
	}
	return GenericUpstream{}
}

func getURL(remote, domain string) (*url.URL, error) {

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

	if domain != "" {
		browserURL.Host = domain
	}

	return &browserURL, nil
}
