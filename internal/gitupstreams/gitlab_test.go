package gitupstreams

import (
	"net/url"
	"testing"
)

func TestGitlabUpstream_BranchURL(t *testing.T) {
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
				repoURL: "https://gitlab.com/user/repo",
				branch:  "master",
			},
			want:    "https://gitlab.com/user/repo",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://gitlab.com/user/project/repo",
				branch:  "feature/demo",
			},
			want:    "https://gitlab.com/user/project/repo/-/tree/feature/demo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := GitlabUpstream{}
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

func TestGitlabUpstream_PullRequestURL(t *testing.T) {
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
				repoURL: "https://gitlab.com/user/project/repo",
				branch:  "master",
			},
			want:    "https://gitlab.com/user/project/repo/-/merge_requests/new",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://gitlab.com/user/project/repo",
				branch:  "feature/demo",
			},
			want:    "https://gitlab.com/user/project/repo/-/merge_requests/new?merge_request%5Bsource_branch%5D=feature%2Fdemo&merge_request%5Btarget_branch%5D=master",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := GitlabUpstream{}
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

func TestGitlabUpstream_CIURL(t *testing.T) {
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
				repoURL: "https://gitlab.com/user/project/repo",
				branch:  "master",
			},
			want:    "https://gitlab.com/user/project/repo/-/pipelines?ref=master&scope=branches",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://gitlab.com/user/project/repo",
				branch:  "feature/demo",
			},
			want:    "https://gitlab.com/user/project/repo/-/pipelines?ref=feature%2Fdemo&scope=branches",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := GitlabUpstream{}
			got, err := u.CIURL(url, tt.args.branch)
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
