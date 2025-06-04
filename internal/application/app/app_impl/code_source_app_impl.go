package app_impl

import (
	"onestep/internal/application/cmd"
	"onestep/internal/common/easytool/util/id_util"
	"onestep/internal/common/enums"
	"onestep/internal/domain/codesource/model"
	"onestep/internal/domain/codesource/service"
	"onestep/internal/domain/codesource/service/service_impl"
	"time"
)

// @Author CY Yan
// @Date 2025/5/30 00:09

// CodeSourceAppImpl is the implementation of CodeSourceApp
type CodeSourceAppImpl struct {
	codeSourceService service.CodeSourceService
}

// NewCodeSourceAppImpl is constructor of CodeSourceAppImpl
func NewCodeSourceAppImpl() *CodeSourceAppImpl {
	return &CodeSourceAppImpl{
		codeSourceService: service_impl.NewCodeSourceServiceImpl(),
	}
}

// Init code source to local, clone remote git repository and create fat, uat, pre, pro branches
func (p *CodeSourceAppImpl) Init(cmd *cmd.CodeSourceCreateCmd) error {
	codeSource := &model.CodeSource{
		Id:                  id_util.GenID(),
		Desc:                cmd.Desc,
		RepositoryType:      "",
		RemoteURL:           cmd.RemoteURL,
		Username:            cmd.Username,
		Password:            cmd.Password,
		PersonalAccessToken: cmd.PersonalAccessToken,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
	// todo bind workspace
	// init code source
	err := p.codeSourceService.Init(codeSource)
	if err != nil {
		panic(err)
	}
	// save code source into persist storage
	return p.codeSourceService.Save(codeSource)
}

// MergeBranchIntoEnv merge branch into env branch
func (p *CodeSourceAppImpl) MergeBranchIntoEnv(mergeCmd *cmd.CodeSourceMergeCmd) {
	codeSource := p.codeSourceService.GetById(mergeCmd.CodeSourceId)
	p.codeSourceService.MergeBranchIntoEnv(codeSource, mergeCmd.Branch, enums.Env(mergeCmd.Env))
}
