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

// CIURL for all generic git repoes
// https://github.com/user/repo/actions?query=branch%3Amain
func (u GenericUpstream) CIURL(repoURL *url.URL, branch string) (string, error) {

	repoURL.Path = repoURL.Path + "/actions"
	q := make(url.Values)
	q.Add("query", "branch:"+branch)
	repoURL.RawQuery = q.Encode()

	return repoURL.String(), nil
}
