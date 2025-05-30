package app_impl

import (
	"onestep/internal/application/cmd"
	"onestep/internal/common/easytool/util/id_util"
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
		Id:             id_util.GenID(),
		Desc:           cmd.Desc,
		RepositoryType: "",
		RemoteURL:      cmd.RemoteURL,
		Username:       cmd.Username,
		Password:       cmd.Password,
		RemoteToken:    cmd.RemoteToken,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	// todo bind workspace
	// init code source
	return p.codeSourceService.Init(codeSource)
}
