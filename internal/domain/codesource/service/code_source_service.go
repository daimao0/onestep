package service

import (
	"onestep/internal/common/enums"
	"onestep/internal/domain/codesource/model"
)

// @Author CY Yan
// @Date 2025/5/28 18:34

// CodeSourceService operation git code source
type CodeSourceService interface {

	// Init the code repository and creates the branches: fat, uat, pre, pro.
	Init(source *model.CodeSource) error

	// Save the code source in persistent storage
	Save(source *model.CodeSource) error

	// GetById get the code source by id
	GetById(id int) *model.CodeSource

	// MergeBranchIntoEnv merge the branch into the environment
	MergeBranchIntoEnv(source *model.CodeSource, branch string, env enums.Env)
}
