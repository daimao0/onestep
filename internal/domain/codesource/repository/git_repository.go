package repository

import (
	"github.com/go-git/go-git/v5"
	"onestep/internal/common/enums"
	"onestep/internal/domain/codesource/model"
)

// @Author CY Yan
// @Date 2025/5/28 18:43

type GitRepository interface {

	// Clone a git branch from remote and return the local path
	// source is the CodeSource domain model
	// branch is the git branch name
	Clone(source *model.CodeSource, branch enums.Env) (string, error)

	// Merge source-branch into target-branch
	Merge(source *model.CodeSource, sourceBranch string, targetBranch string) error

	// Checkout a git branch
	Checkout(source *model.CodeSource, branch string) (*git.Worktree, error)

	// Pull a git branch
	Pull(source *model.CodeSource, branch string) error

	// CheckBranchExist  a git branch
	CheckBranchExist(source *model.CodeSource, branch enums.Env) (bool, error)

	// GetGitRepository  get git repository to operate git
	GetGitRepository(source *model.CodeSource) (*git.Repository, error)

	// BranchFromMaster  branch from master
	BranchFromMaster(source *model.CodeSource, branch enums.Env) error
}
