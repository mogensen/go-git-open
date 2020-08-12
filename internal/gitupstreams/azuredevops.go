package gitupstreams

import (
	"net/url"
	"strings"
)

// AzureDevopsUpstream is tested for
// - dev.azure.com
// - ssh.dev.azure.com
type AzureDevopsUpstream struct{}

// WillHandle for azure devops
func (u AzureDevopsUpstream) WillHandle(repoURL *url.URL) bool {
	return strings.Contains(strings.ToLower(repoURL.Host), "azure.com")
}

// BranchURL creates a browser url for Azure DevOps
// https://ssh.dev.azure.com/v3/CORP/Project/GitRepo
// https://dev.azure.com/CORP/Project/_git/GitRepo
// For branch:
// https://dev.azure.com/CORP/Project/_git/GitRepo?version=GBdevelop
func (u AzureDevopsUpstream) BranchURL(repoURL *url.URL, branch string) (string, error) {

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

	return repoURL.String(), nil
}
