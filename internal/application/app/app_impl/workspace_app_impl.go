package app_impl

import (
	"onestep/internal/application/cmd"
	"onestep/internal/application/dto"
	"onestep/internal/domain/workspace/model/factory"
	"onestep/internal/domain/workspace/model/factory/factory_impl"
	"onestep/internal/domain/workspace/service"
	"onestep/internal/domain/workspace/service/service_impl"
)

type WorkspaceAppImpl struct {
	service.WorkspaceService
	factory.WorkspaceFactory
}

func NewWorkspaceAppImpl() *WorkspaceAppImpl {
	return &WorkspaceAppImpl{
		service_impl.NewWorkspaceServiceImpl(),
		factory_impl.NewWorkspaceFactoryImpl(),
	}
}

// Create creates a new workspace
func (p *WorkspaceAppImpl) Create(cmd *cmd.WorkspaceCreatCmd) error {
	// create a workspace by workspace factory
	workspace := p.WorkspaceFactory.CreateWorkspace(cmd.Name)
	return p.WorkspaceService.Create(workspace)
}

// Update updates a workspace by ID
func (p *WorkspaceAppImpl) Update(cmd *cmd.WorkspaceUpdateCmd) error {
	return nil

}

// GetById gets a workspace by ID
func (p *WorkspaceAppImpl) GetById(id int) (*dto.WorkspaceDTO, error) {

	return nil, nil
}

// RemoveById removes a workspace by ID
func (p *WorkspaceAppImpl) RemoveById(id int) error {

	return nil
}
