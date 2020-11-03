package gitupstreams

import (
	"net/url"
	"strings"
)

// BitbucketOrgUpstream is tested for
// - Bitbucket.org
type BitbucketOrgUpstream struct{}

// WillHandle for generic is always true.. This is used as a sane fallback
func (u BitbucketOrgUpstream) WillHandle(repoURL *url.URL) bool {
	return strings.Contains(strings.ToLower(repoURL.Host), "bitbucket.org")
}

// BranchURL creates a browser url for Bitbucket.Org
// For branch:
//  - https://bitbucket.org/fdfapps/raceapp/src/HEAD/?at=feature%2Fdemo
func (u BitbucketOrgUpstream) BranchURL(repoURL *url.URL, branch string) (string, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/src/HEAD/"
		q := make(url.Values)
		q.Add("at", branch)
		repoURL.RawQuery = q.Encode()
	}
	return repoURL.String(), nil
}

// PullRequestURL creates a browser url for Bitbucket.Org
// For branch:
//  - https://bitbucket.org/fdfapps/raceapp/pull-requests/new?source=feature%2FCaptureCheckpoints
func (u BitbucketOrgUpstream) PullRequestURL(repoURL *url.URL, branch string) (string, error) {
	repoURL.Path = repoURL.Path + "/pull-requests/new"
	q := make(url.Values)
	q.Add("source", branch)
	repoURL.RawQuery = q.Encode()
	return repoURL.String(), nil
}

// CIURL creates a browser url for Bitbucket.Org
// For branch:
//  - https://bitbucket.org/fdfapps/raceapp/addon/pipelines/home
func (u BitbucketOrgUpstream) CIURL(repoURL *url.URL, branch string) (string, error) {
	repoURL.Path = repoURL.Path + "/addon/pipelines/home"
		if branch != "master" {
		repoURL.Path = repoURL.Path + "/src/HEAD/"
		q := make(url.Values)
		q.Add("at", branch)
		repoURL.RawQuery = q.Encode()
	}
	return repoURL.String(), nil
}
