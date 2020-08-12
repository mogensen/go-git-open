package main

import (
	"log"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/format/config"
	"github.com/go-git/go-git/v5/storage/memory"
)

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

func Test_getOverwriteDomain(t *testing.T) {
	type args struct {
		gitRemote  string
		openDomain string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				gitRemote:  "git@github.com:git-fixtures/basic.git",
				openDomain: "myrepo.com",
			},
			want: "myrepo.com",
		},
		{
			name: "simple",
			args: args{
				gitRemote:  "git@github.com:git-fixtures/basic.git",
				openDomain: "",
			},
			want: "",
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
				ReferenceName: plumbing.NewBranchReferenceName("master"),
			})

			c, _ := gitRepo.Config()
			c.Raw.Sections = append(c.Raw.Sections, &config.Section{Name: "open", Options: config.Options{&config.Option{Key: "domain", Value: tt.args.openDomain}}})

			if err != nil {
				log.Fatal(err)
			}

			if got := getOverwriteDomain(gitRepo); got != tt.want {
				t.Errorf("getOverwriteDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
