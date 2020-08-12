package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
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

			dir, err := ioutil.TempDir("", "clone-example-"+tt.name)
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir) // clean up

			gitRepo, err := newRepo(dir, tt.args.gitRemote, tt.args.gitBranch)
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

			dir, err := ioutil.TempDir("", "clone-example-"+tt.name)
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
