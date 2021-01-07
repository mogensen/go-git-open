package gitupstreams

import (
	"net/url"
	"strings"
)

// GitlabUpstream is tested for
// - gitlab.com
type GitlabUpstream struct{}

// WillHandle for generic is always true.. This is used as a sane fallback
func (u GitlabUpstream) WillHandle(repoURL *url.URL) bool {
	return strings.Contains(strings.ToLower(repoURL.Host), "gitlab")
}

// BranchURL creates a browser url for gitlab.com
// For branch:
//  - https://gitlab.com/fdfapps/core/raceapp/-/tree/feature/CaptureCheckpoints
func (u GitlabUpstream) BranchURL(repoURL *url.URL, branch string) (string, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/-/tree/" + branch
	}
	return repoURL.String(), nil
}

// PullRequestURL creates a browser url for gitlab.com
// For branch:
//  - https://gitlab.com/fdfapps/core/repo/-/merge_requests/new?merge_request[source_branch]=feature/CaptureCheckpoints&merge_request[target_branch]=master
func (u GitlabUpstream) PullRequestURL(repoURL *url.URL, branch string) (string, error) {
	repoURL.Path = repoURL.Path + "/-/merge_requests/new"
	if branch != "master" {
		q := make(url.Values)
		q.Add("merge_request[source_branch]", branch)
		q.Add("merge_request[target_branch]", "master")

		repoURL.RawQuery = q.Encode()
	}
	return repoURL.String(), nil
}

// CIURL creates a browser url for gitlab.com
// For branch:
//  - https://gitlab.com/fdfapps/core/raceapp/-/pipelines
//  - https://gitlab.com/fdfapps/core/raceapp/-/pipelines?scope=branches&ref=feature/CaptureCheckpoints
func (u GitlabUpstream) CIURL(repoURL *url.URL, branch string) (string, error) {
	repoURL.Path = repoURL.Path + "/-/pipelines"
	q := make(url.Values)
	q.Add("scope", "branches")
	q.Add("ref", branch)
	repoURL.RawQuery = q.Encode()
	return repoURL.String(), nil
}
