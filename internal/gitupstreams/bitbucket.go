package gitupstreams

import (
	"net/url"
	"strings"
)

// BitbucketUpstream is tested for
// - Bitbucket
type BitbucketUpstream struct{}

// WillHandle for generic is always true.. This is used as a sane fallback
func (u BitbucketUpstream) WillHandle(repoURL *url.URL) bool {
	return strings.Contains(strings.ToLower(repoURL.Host), "bitbucket")
}

// BranchURL creates a browser url for Bitbucket
// For branch:

//  - https://bitbucket.mydomain.com/projects/fdfapps/repos/raceapp/browse?at=refs%2Fheads%2Ffeature%2Fdemo
func (u BitbucketUpstream) BranchURL(repoURL *url.URL, branch string) (string, error) {
	u.cleanURL(repoURL)
	repoURL.Path = repoURL.Path + "/browse"
	if branch != "master" {
		q := make(url.Values)
		q.Add("at", "refs/heads/"+branch)
		repoURL.RawQuery = q.Encode()
	}
	return repoURL.String(), nil
}

// PullRequestURL creates a browser url for Bitbucket
// For branch:
//  - https://bitbucket.mydomain.com/projects/fdfapps/repos/raceapp/pull-requests?create&sourceBranch=refs%2Fheads%2Ffeature%2FCaptureCheckpoints
func (u BitbucketUpstream) PullRequestURL(repoURL *url.URL, branch string) (string, error) {
	u.cleanURL(repoURL)
	repoURL.Path = repoURL.Path + "/pull-requests"
	q := make(url.Values)
	q.Add("create", "")
	q.Add("sourceBranch", "refs/heads/"+branch)
	repoURL.RawQuery = q.Encode()
	return repoURL.String(), nil
}

// CIURL creates a browser url for Bitbucket
// For branch:
//  - https://bitbucket.mydomain.com/fdfapps/raceapp/addon/pipelines/home
func (u BitbucketUpstream) CIURL(repoURL *url.URL, branch string) (string, error) {
	u.cleanURL(repoURL)
	repoURL.Path = repoURL.Path + "/builds"
	if branch != "master" {
		q := make(url.Values)
		q.Add("branch", "refs/heads/"+branch)
		repoURL.RawQuery = q.Encode()
	}
	return repoURL.String(), nil
}

func (u BitbucketUpstream) cleanURL(repoURL *url.URL) {
	pathParts := strings.Split(repoURL.Path, "/")
	newParts := []string{}
	for _, part := range pathParts {
		if part != "scm" && part != "" {
			newParts = append(newParts, part)
		}
	}

	repoURL.Path = "projects/" + newParts[0] + "/repos/" + newParts[1]
}
