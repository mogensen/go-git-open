package gitupstreams

import (
	"net/url"
	"testing"
)

func TestBitbucketOrgUpstream_BranchURL(t *testing.T) {
	type args struct {
		repoURL string
		branch  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "https on master",
			args: args{
				repoURL: "https://bitbucket.org/user/repo",
				branch:  "master",
			},
			want:    "https://bitbucket.org/user/repo",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://bitbucket.org/user/repo",
				branch:  "feature/demo",
			},
			want:    "https://bitbucket.org/user/repo/src/HEAD/?at=feature%2Fdemo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := BitbucketOrgUpstream{}
			got, err := u.BranchURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("AzureURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AzureURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitbucketOrgUpstream_PullRequestURL(t *testing.T) {
	type args struct {
		repoURL string
		branch  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "https on master",
			args: args{
				repoURL: "https://bitbucket.org/user/repo",
				branch:  "master",
			},
			want:    "https://bitbucket.org/user/repo/pull-requests/new?source=master",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://bitbucket.org/user/repo",
				branch:  "feature/demo",
			},
			want:    "https://bitbucket.org/user/repo/pull-requests/new?source=feature%2Fdemo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := BitbucketOrgUpstream{}
			got, err := u.PullRequestURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("AzureURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AzureURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
