package gitupstreams

import (
	"testing"
)

func TestGenericUpstream_BranchURL(t *testing.T) {
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
				repoURL: "https://github.com/user/repo",
				branch:  "master",
			},
			want:    "https://github.com/user/repo",
			wantErr: false,
		},
		{
			name: "https on develop",
			args: args{
				repoURL: "https://github.com/user/repo",
				branch:  "develop",
			},
			want:    "https://github.com/user/repo/tree/develop",
			wantErr: false,
		},
		{
			name: "git on master",
			args: args{
				repoURL: "git@github.com:user/repo.git",
				branch:  "master",
			},
			want:    "https://github.com/user/repo",
			wantErr: false,
		},
		{
			name: "git on develop",
			args: args{
				repoURL: "git@github.com:user/repo.git",
				branch:  "develop",
			},
			want:    "https://github.com/user/repo/tree/develop",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := getURL(tt.args.repoURL, "")
			u := GenericUpstream{}
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

func TestGenericUpstream_PullRequestURL(t *testing.T) {
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
				repoURL: "https://github.com/user/repo",
				branch:  "master",
			},
			want:    "https://github.com/user/repo",
			wantErr: false,
		},
		{
			name: "https on develop",
			args: args{
				repoURL: "https://github.com/user/repo",
				branch:  "develop",
			},
			want:    "https://github.com/user/repo/compare/master...develop",
			wantErr: false,
		},
		{
			name: "git on master",
			args: args{
				repoURL: "git@github.com:user/repo.git",
				branch:  "master",
			},
			want:    "https://github.com/user/repo",
			wantErr: false,
		},
		{
			name: "git on develop",
			args: args{
				repoURL: "git@github.com:user/repo.git",
				branch:  "develop",
			},
			want:    "https://github.com/user/repo/compare/master...develop",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url, _ := getURL(tt.args.repoURL, "")
			u := GenericUpstream{}
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
