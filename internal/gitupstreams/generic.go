package gitupstreams

import (
	"net/url"
)

// GenericUpstream is tested for
// - Github.com
// - Gist.Github.com
type GenericUpstream struct{}

// WillHandle for generic is always true.. This is used as a sane fallback
func (u GenericUpstream) WillHandle(repoURL *url.URL) bool {
	return true
}

// BranchURL for all generic git repoes
func (u GenericUpstream) BranchURL(repoURL *url.URL, branch string) (string, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/tree/" + branch
	}
	return repoURL.String(), nil
}

// PullRequestURL for all generic git repoes
// https://github.com/user/repo/compare/master...develop
func (u GenericUpstream) PullRequestURL(repoURL *url.URL, branch string) (string, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/compare/master..." + branch
	}
	return repoURL.String(), nil
}
