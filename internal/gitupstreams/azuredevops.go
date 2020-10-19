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

	u.cleanURL(repoURL)
	if branch != "master" {
		q := make(url.Values)
		q.Add("version", "GB"+branch)
		repoURL.RawQuery = q.Encode()
	}

	return repoURL.String(), nil
}

// PullRequestURL creates a browser url for Azure DevOps
// https://ssh.dev.azure.com/v3/CORP/Project/GitRepo
// https://dev.azure.com/CORP/Project/_git/GitRepo
// For pull-requests:
// https://dev.azure.com/CORP/Project/_git/GitRepo/pullrequestcreate?sourceRef=develop&targetRef=master
func (u AzureDevopsUpstream) PullRequestURL(repoURL *url.URL, branch string) (string, error) {

	u.cleanURL(repoURL)
	repoURL.Path += "/pullrequestcreate"

	if branch != "master" {
		q := make(url.Values)
		q.Add("sourceRef", branch)
		q.Add("targetRef", "master")
		repoURL.RawQuery = q.Encode()
	}

	return repoURL.String(), nil
}

func (u AzureDevopsUpstream) cleanURL(repoURL *url.URL) {
	pathParts := strings.Split(repoURL.Path, "/")
	newParts := []string{}
	for _, part := range pathParts {
		if part != "_git" && part != "v3" && part != "" {
			newParts = append(newParts, part)
		}
	}

	repoURL.Path = newParts[0] + "/" + newParts[1] + "/_git/" + newParts[2]

	if repoURL.Host == "ssh.dev.azure.com" {
		repoURL.Host = "dev.azure.com"
	}
}
