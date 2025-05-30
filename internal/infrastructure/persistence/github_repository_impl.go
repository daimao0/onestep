package persistence

import (
	"context"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"onestep/internal/common/enums"
	"onestep/internal/domain/codesource/model"
)

// @Author CY Yan
// @Date 2025/5/28 18:54

// GithubRepositoryImpl to operation remote git repository
type GithubRepositoryImpl struct {
}

// NewGithubRepositoryImpl to create a new GithubRepositoryImpl
func NewGithubRepositoryImpl() *GithubRepositoryImpl {
	return &GithubRepositoryImpl{}
}

// Clone a git branch from remote and return the local path
// source is the CodeSource domain model
// branch is the git branch name
func (p *GithubRepositoryImpl) Clone(source *model.CodeSource, branch enums.Env) (string, error) {
	options := &git.CloneOptions{
		URL:           source.RemoteURL,
		Auth:          source.GetGitBasicAuth(),
		ReferenceName: plumbing.NewBranchReferenceName(branch.String()),
	}
	_, err := git.PlainClone(source.GetGitRepoDirPath()+"_"+branch.String(), false, options)
	if err != nil {
		log.Println(err.Error())
	}
	return source.GetProjectName(), err
}

// CheckBranchExist  check if the 'env' branch exist
// if the function returns error_code != nil, you cannot determine whether the branch exists.
func (p *GithubRepositoryImpl) CheckBranchExist(source *model.CodeSource, env enums.Env) (bool, error) {
	// get git repository to operate git
	repo, err := p.GetGitRepository(source)
	if err != nil {
		return false, err
	}

	// find the env branch in local git repository
	_, err = repo.Reference(plumbing.NewBranchReferenceName(env.String()), true)
	if err == nil {
		return true, err
	}

	// if local branch not exist, check if the remote branch exist
	remote, err := repo.Remote("origin")
	if err != nil {
		return false, err
	}

	// list remote branches
	rfs, err := remote.List(&git.ListOptions{Auth: source.GetGitBasicAuth()})
	if err != nil {
		return false, err
	}

	//  check if the remote branch exist env branch
	for _, ref := range rfs {
		if ref.Name().String() == "refs/heads/"+env.String() {
			// git fetch a branch
			refSpec := fmt.Sprintf("+refs/heads/%s:refs/remotes/origin/%s", env, env)
			_ = repo.Fetch(&git.FetchOptions{
				RemoteName: "origin",
				RefSpecs:   []config.RefSpec{config.RefSpec(refSpec)},
				Auth:       source.GetGitBasicAuth(),
				Depth:      1,
			})
			worktree, err := repo.Worktree()
			if err != nil {
				return false, nil
			}
			// checkout the remote branch
			err = worktree.Checkout(&git.CheckoutOptions{
				Branch: plumbing.NewBranchReferenceName(env.String()),
				Create: true,
			})
			if err != nil {
				return false, nil
			}
			return true, nil
		}
	}
	//env branch not exist in local or remote
	return false, nil
}

// GetGitRepository to operate git repository head: master
func (p *GithubRepositoryImpl) GetGitRepository(source *model.CodeSource) (*git.Repository, error) {
	// open git repository on local
	open, _ := git.PlainOpen(source.GetGitRepoDirPath())
	if open != nil {
		return open, nil
	}
	// clone git repository from remote if not exist
	options := &git.CloneOptions{
		URL:  source.RemoteURL,
		Auth: source.GetGitBasicAuth(),
	}
	repo, err := git.PlainCloneContext(context.Background(), source.GetGitRepoDirPath(), false, options)
	return repo, err
}

// BranchFromMaster  branch from master
func (p *GithubRepositoryImpl) BranchFromMaster(source *model.CodeSource, env enums.Env) error {
	// get git repository to operate git
	repo, err := p.GetGitRepository(source)
	if err != nil {
		return err
	}
	ref, err := repo.Reference(plumbing.Master, true)
	if err != nil {
		return err
	}
	// create a new env branch for master
	newRefName := plumbing.NewBranchReferenceName(env.String())
	newRef := plumbing.NewHashReference(newRefName, ref.Hash())

	// set reference in Store
	err = repo.Storer.SetReference(newRef)
	if err != nil {
		return err
	}

	//push env branch to remote
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       source.GetGitBasicAuth(),
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%s", env, env)),
		},
	})
	return err
}
