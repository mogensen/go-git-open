package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func Test_getBrowerURL(t *testing.T) {
	tests := []struct {
		name            string
		branch          string
		remoteURL       string
		domainOverwrite string
		want            string
		wantErr         bool
	}{
		{
			name:      "gh: basic",
			branch:    "master",
			remoteURL: "git@github.com:user/repo.git",
			want:      "https://github.com/user/repo",
			wantErr:   false,
		},
		{
			name:      "gh: basic with branch",
			branch:    "develop",
			remoteURL: "git@github.com:user/repo.git",
			want:      "https://github.com/user/repo/tree/develop",
			wantErr:   false,
		},
		{
			name:      "gh: basic http",
			branch:    "master",
			remoteURL: "http://github.com/user/repo.git",
			want:      "http://github.com/user/repo",
			wantErr:   false,
		},
		{
			name:            "gh: basic with domain overwrite",
			branch:          "master",
			remoteURL:       "git@github.com:user/repo.git",
			domainOverwrite: "repo.git.com",
			want:            "https://repo.git.com/user/repo",
			wantErr:         false,
		},
		{
			name:      "azure devops: basic",
			branch:    "master",
			remoteURL: "https://CORP@dev.azure.com/v3/CORP/Project/GitRepo",
			want:      "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr:   false,
		},
		{
			name:      "azure devops: ssh",
			branch:    "master",
			remoteURL: "git@ssh.dev.azure.com:v3/CORP/Project/GitRepo",
			want:      "https://dev.azure.com/CORP/Project/_git/GitRepo",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBrowerURL(tt.remoteURL, tt.domainOverwrite, tt.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBrowerURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBrowerURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURLFromGitRepo(t *testing.T) {
	type args struct {
		gitRemote string
		gitBranch string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "gh: basic",
			args: args{
				gitRemote: "git@github.com:git-fixtures/basic.git",
				gitBranch: "master",
			},
			want:    "https://github.com/git-fixtures/basic",
			wantErr: false,
		},
		{
			name: "gh: basic",
			args: args{
				gitRemote: "git@github.com:git-fixtures/basic.git",
				gitBranch: "branch",
			},
			want:    "https://github.com/git-fixtures/basic/tree/branch",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fs := memfs.New()
			// Git objects storer based on memory
			storer := memory.NewStorage()

			// Clones the repository into the worktree (fs) and storer all the .git
			// content into the storer
			gitRepo, err := git.Clone(storer, fs, &git.CloneOptions{
				URL:           tt.args.gitRemote,
				SingleBranch:  true,
				ReferenceName: plumbing.NewBranchReferenceName(tt.args.gitBranch),
			})
			if err != nil {
				log.Fatal(err)
			}

			got, err := getURLFromGitRepo(gitRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLFromGitRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getURLFromGitRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
