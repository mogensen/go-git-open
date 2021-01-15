package gitupstreams

import (
	"net/url"
	"testing"
)

func TestBitbucketUpstream_BranchURL(t *testing.T) {
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
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "master",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/browse?at=refs%2Fheads%2Fmaster",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "feature/demo",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/browse?at=refs%2Fheads%2Ffeature%2Fdemo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := BitbucketUpstream{}
			got, err := u.BranchURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("BranchURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BranchURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitbucketUpstream_PullRequestURL(t *testing.T) {
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
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "master",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/pull-requests?create=true&sourceBranch=refs%2Fheads%2Fmaster",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "feature/demo",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/pull-requests?create=true&sourceBranch=refs%2Fheads%2Ffeature%2Fdemo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := BitbucketUpstream{}
			got, err := u.PullRequestURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullRequestURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PullRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitbucketUpstream_CIURL(t *testing.T) {
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
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "master",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/builds?branch=refs%2Fheads%2Fmaster",
			wantErr: false,
		},
		{
			name: "https on feature/demo",
			args: args{
				repoURL: "https://bitbucket.example.com/scm/user/repo",
				branch:  "feature/demo",
			},
			want:    "https://bitbucket.example.com/projects/user/repos/repo/builds?branch=refs%2Fheads%2Ffeature%2Fdemo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := url.Parse(tt.args.repoURL)
			u := BitbucketUpstream{}
			got, err := u.CIURL(url, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("CIURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CIURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
