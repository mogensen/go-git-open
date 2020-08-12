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
