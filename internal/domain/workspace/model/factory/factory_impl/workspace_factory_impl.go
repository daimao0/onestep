package factory_impl

import (
	"onestep/internal/common/easytool/util/id_util"
	"onestep/internal/domain/workspace/model"
	"time"
)

type WorkspaceFactoryImpl struct {
}

func NewWorkspaceFactoryImpl() *WorkspaceFactoryImpl {
	return &WorkspaceFactoryImpl{}
}

// CreateWorkspace is the factory method to create a new workspace
func (p *WorkspaceFactoryImpl) CreateWorkspace(name string) *model.Workspace {
	return &model.Workspace{
		Id:        int(id_util.GenID()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
