package repository

import "onestep/internal/domain/codesource/model"

// @Author CY Yan
// @Date 2025/5/28 18:43

type GitRepository interface {

	// Clone a git branch from remote and return the local path
	Clone(source *model.CodeSource) (string, error)
}
