package service_impl

import (
	"onestep/internal/common/enums"
	"onestep/internal/common/error_code"
	"onestep/internal/domain/codesource/model"
	"onestep/internal/domain/codesource/repository"
	"onestep/internal/infrastructure/persistence"
)

// @Author CY Yan
// @Date 2025/5/28 23:24

// CodeSourceServiceImpl is the implementation of CodeSourceService
type CodeSourceServiceImpl struct {
	gitRepository repository.GitRepository
}

// NewCodeSourceServiceImpl is the constructor of CodeSourceServiceImpl
func NewCodeSourceServiceImpl() *CodeSourceServiceImpl {
	return &CodeSourceServiceImpl{
		gitRepository: persistence.NewGithubRepositoryImpl(),
	}
}

// Init the code repository and creates the branches: fat, uat, pre, pro.
func (p *CodeSourceServiceImpl) Init(source *model.CodeSource) error {

	exist, err := p.gitRepository.CheckBranchExist(source, enums.FAT)
	if !exist {
		err = p.gitRepository.BranchFromMaster(source, enums.FAT)
	}

	exist, _ = p.gitRepository.CheckBranchExist(source, enums.UAT)
	if !exist {
		err = p.gitRepository.BranchFromMaster(source, enums.UAT)
	}

	exist, _ = p.gitRepository.CheckBranchExist(source, enums.PRE)
	if !exist {
		err = p.gitRepository.BranchFromMaster(source, enums.PRE)
	}

	exist, _ = p.gitRepository.CheckBranchExist(source, enums.PRO)
	if !exist {
		err = p.gitRepository.BranchFromMaster(source, enums.PRO)
	}

	if err != nil {
		return error_code.NewErrorCode(error_code.CodeSourceInitFailed, err)
	}

	return nil
}
