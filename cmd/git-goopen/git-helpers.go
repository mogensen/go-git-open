package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"

	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Helpers is based on
// - https://github.com/thecjharries/go-git-ref-bug/

func newRepo(directory, repoURL, branch string) (*git.Repository, error) {

	r, err := git.PlainInit(directory, false)
	if git.ErrRepositoryAlreadyExists == err {
		r, err = git.PlainOpen(directory)
	}
	if err != nil {
		return nil, err
	}
	makeACommit(r, directory, "master")
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "example",
		URLs: []string{"https://github.com/git-fixtures/basic.git"},
	})
	if branch != "master" {
		createNewBranch(r, branch)
	}
	return r, nil
}

func makeACommit(repo *git.Repository, directory string, random_stuff string) {
	// Pulled from the commit demo
	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	// Info("echo \"hello world!\" > example-git-file")
	filename := filepath.Join(directory, "example-git-file")
	err = ioutil.WriteFile(filename, []byte(random_stuff), 0644)
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	// Info("git add example-git-file")
	_, err = w.Add("example-git-file")
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	// Info("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	fmt.Println(status)
	// Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	// Info("git show -s")
	obj, err := repo.CommitObject(commit)
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	fmt.Println(obj)
}

func createNewBranch(repo *git.Repository, branchName string) {
	// Pulled from the branch demo
	repoConfig, err := repo.Config()
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	newRef := plumbing.NewHashReference(plumbing.NewBranchReferenceName(branchName), headRef.Hash())

	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: newRef.Name(),
		Keep:   true,
		Create: true,
	})
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	err = repo.Storer.SetReference(newRef)
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	err = repo.Storer.SetConfig(repoConfig)
	if err != nil {
		log.Fatalf("Broke: %v", err)
	}
	makeACommit(repo, w.Filesystem.Root(), branchName)
}

func saveGitConfig(directory string, conf *config.Config) error {
	serializedConfig, err := conf.Marshal()
	if err != nil {
		return fmt.Errorf("cannot serialize git configuration. Error: %w", err)
	}
	err = ioutil.WriteFile(path.Join(directory, ".git", "config"), serializedConfig, 0644)
	return err
}
