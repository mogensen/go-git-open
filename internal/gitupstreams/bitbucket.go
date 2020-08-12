package gitupstreams

import (
	"net/url"
)

// BitbucketOrgURL creates a browser url for:
// - Bitbucket.org
// with branch: https://bitbucket.org/fdfapps/raceapp/src/HEAD/?at=feature%2Fdemo
func BitbucketOrgURL(repoURL *url.URL, branch string) (*url.URL, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/src/HEAD/"
		q := make(url.Values)
		q.Add("at", branch)
		repoURL.RawQuery = q.Encode()
	}
	return repoURL, nil
}
