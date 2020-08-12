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
