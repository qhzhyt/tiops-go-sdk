package utils

import (
	"errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"time"
)

func OpenGitRepo(repoPath, worktreePath string) (*git.Repository, error) {
	repoFS := osfs.New(repoPath)
	wtFS := osfs.New(worktreePath)
	store := filesystem.NewStorage(repoFS, cache.NewObjectLRUDefault())
	return git.Open(store, wtFS)
}

func GitInit(repoPath string) error {
	//repoPath := path.Join(shared.Config.DataPath, "git-repos", gitRepo)
	//worktreePath := path.Join(shared.Config.ConfigPath, gitRepo)
	repoFS := osfs.New(repoPath)
	//wtFS := osfs.New(worktreePath)
	store := filesystem.NewStorage(repoFS, cache.NewObjectLRUDefault())
	_, err := git.Init(store, nil)

	return err
	//git.PlainInit()
}

func GitCommit(repoPath, worktreePath, msg, author, email string) error {
	var err error
	var repo *git.Repository
	var worktree *git.Worktree
	var hash plumbing.Hash
	repo, err = OpenGitRepo(repoPath, worktreePath)
	if err != nil {
		return err
	}
	worktree, err = repo.Worktree()
	if err != nil {
		return err
	}

	var status git.Status
	status, err = worktree.Status()
	if err != nil {
		return err
	}

	if status.IsClean() {
		return errors.New("worktree is clean, no files to commit")
	}

	_, err = worktree.Add(".")
	if err != nil {
		return err
	}
	hash, err = worktree.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{Name: author, Email: email, When: time.Now()},
	})
	if err != nil {
		return err
	}
	//var commit *object.Commit
	_, err = repo.CommitObject(hash)
	//fmt.Println(commit)
	return err
}

type GitLogItem struct {
	ID      string    `json:"id"`
	Author  string    `json:"author"`
	Email   string    `json:"email"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func GitLog(repoPath, worktreePath string) ([]*GitLogItem, error) {
	var err error
	var repo *git.Repository
	repo, err = OpenGitRepo(repoPath, worktreePath)
	if err != nil {
		return nil, err
	}
	var ref *plumbing.Reference
	ref, err = repo.Head()
	if err != nil {
		return nil, err
	}
	var commits object.CommitIter
	// ... retrieves the commit history
	commits, err = repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	var result []*GitLogItem
	err = commits.ForEach(func(commit *object.Commit) error {
		result = append(result, &GitLogItem{
			ID:      commit.ID().String(),
			Author:  commit.Author.Name,
			Email:   commit.Author.Email,
			Time:    commit.Author.When,
			Message: commit.Message,
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

type GitStatusItem struct {
	FilePath string `json:"filePath"`
	Status   string `json:"status"`
	Extra    string `json:"extra"`
	Staging  string `json:"staging"`
}

func GitStatus(repoPath, worktreePath string) ([]*GitStatusItem, error) {
	var err error
	var repo *git.Repository
	var worktree *git.Worktree
	repo, err = OpenGitRepo(repoPath, worktreePath)
	if err != nil {
		return nil, err
	}
	worktree, err = repo.Worktree()
	if err != nil {
		return nil, err
	}
	var status git.Status
	status, err = worktree.Status()
	if err != nil {
		return nil, err
	}

	var result []*GitStatusItem

	for filePath, fileStatus := range status {
		if fileStatus.Worktree != git.Unmodified {
			result = append(result, &GitStatusItem{FilePath: filePath, Status: string(fileStatus.Worktree), Extra: fileStatus.Extra, Staging: string(fileStatus.Staging)})
		}
	}

	return result, nil
}

func GitCheckout(repoPath, worktreePath string) error {
	var err error
	var repo *git.Repository
	var worktree *git.Worktree
	repo, err = OpenGitRepo(repoPath, worktreePath)
	if err != nil {
		return err
	}
	worktree, err = repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Checkout(&git.CheckoutOptions{Force: true})
	if err != nil {
		return err
	}

	return nil
}
