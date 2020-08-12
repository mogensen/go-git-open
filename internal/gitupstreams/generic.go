package gitupstreams

import (
	"net/url"
)

// GenericURL creates a browser url for:
// - Github
func GenericURL(repoURL *url.URL, branch string) (*url.URL, error) {
	if branch != "master" {
		repoURL.Path = repoURL.Path + "/tree/" + branch
	}
	return repoURL, nil
}
