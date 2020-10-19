package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

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
			name: "simple with overwrite",
			args: args{
				gitRemote:  "https://github.com/git-fixtures/basic",
				openDomain: "myrepo.com",
			},
			want: "myrepo.com",
		},
		{
			name: "simple without overwrite",
			args: args{
				gitRemote:  "https://github.com/git-fixtures/basic",
				openDomain: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir, err := ioutil.TempDir("", "go-git-open")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir) // clean up

			gitRepo, err := newRepo(dir, tt.args.gitRemote, "master")
			if err != nil {
				log.Fatal(err)
			}

			// Add git config open.domain
			c, _ := gitRepo.Config()
			c.Raw.AddOption("open", "", "domain", tt.args.openDomain)
			saveGitConfig(dir, c)

			if err != nil {
				log.Fatal(err)
			}

			if got := getOverwriteDomain(gitRepo); got != tt.want {
				t.Errorf("getOverwriteDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRepoInfo(t *testing.T) {
	type args struct {
		gitRemote string
		gitBranch string
	}
	tests := []struct {
		name       string
		args       args
		wantRemote string
		wantDomain string
		wantBranch string
		wantTag    string
		wantErr    bool
	}{
		{
			name: "gh: basic",
			args: args{
				gitRemote: "git@github.com:git-fixtures/basic.git",
				gitBranch: "master",
			},
			wantRemote: "https://github.com/git-fixtures/basic.git",
			wantBranch: "master",
			wantErr:    false,
		},
		{
			name: "gh: basic",
			args: args{
				gitRemote: "git@github.com:git-fixtures/basic.git",
				gitBranch: "branch",
			},
			wantRemote: "https://github.com/git-fixtures/basic.git",
			wantBranch: "branch",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir, err := ioutil.TempDir("", "go-git-open")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir) // clean up

			gitRepo, err := newRepo(dir, tt.args.gitRemote, tt.args.gitBranch)
			if err != nil {
				log.Fatal(err)
			}

			gotRemote, gotDomain, gotBranch, gotTag, err := getRepoInfo(gitRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRepoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRemote != tt.wantRemote {
				t.Errorf("getRepoInfo() gotRemote = %v, want %v", gotRemote, tt.wantRemote)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("getRepoInfo() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
			if gotBranch != tt.wantBranch {
				t.Errorf("getRepoInfo() gotBranch = %v, want %v", gotBranch, tt.wantBranch)
			}
			if gotTag != tt.wantTag {
				t.Errorf("getRepoInfo() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
		})
	}
}
