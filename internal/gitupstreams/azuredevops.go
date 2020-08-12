package gitupstreams

import (
	"net/url"
	"strings"
)

// AzureURL creates a browser url for Azure DevOps
// https://ssh.dev.azure.com/v3/CORP/Project/GitRepo
// https://dev.azure.com/CORP/Project/_git/GitRepo
// For branch:
// https://dev.azure.com/CORP/Project/_git/GitRepo?version=GBdevelop
func AzureURL(repoURL *url.URL, branch string) (*url.URL, error) {

	repoURL.Path = strings.TrimPrefix(repoURL.Path, "v3/")
	repoURL.Path = strings.TrimPrefix(repoURL.Path, "/v3/")

	pathParts := strings.Split(repoURL.Path, "/")
	repoURL.Path = pathParts[0] + "/" + pathParts[1] + "/_git/" + pathParts[2]

	if repoURL.Host == "ssh.dev.azure.com" {
		repoURL.Host = "dev.azure.com"
	}

	if branch != "master" {
		q := make(url.Values)
		q.Add("version", "GB"+branch)
		repoURL.RawQuery = q.Encode()
	}

	return repoURL, nil
}
