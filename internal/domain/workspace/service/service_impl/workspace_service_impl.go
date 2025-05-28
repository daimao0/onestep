package service_impl

import (
	"onestep/internal/domain/workspace/model"
	"onestep/internal/domain/workspace/repository"
	"onestep/internal/infrastructure/persistence"
)

// WorkspaceServiceImpl is the implementation of workspace domain Service
type WorkspaceServiceImpl struct {
	WorkspaceRepository repository.WorkspaceRepository
}

// NewWorkspaceServiceImpl is the constructor of workspaceServiceImpl
func NewWorkspaceServiceImpl() *WorkspaceServiceImpl {
	return &WorkspaceServiceImpl{
		WorkspaceRepository: persistence.NewWorkspaceRepositoryImpl(),
	}
}

// Create a new workspace
func (p *WorkspaceServiceImpl) Create(workspace *model.Workspace) error {
	err := p.WorkspaceRepository.Insert(workspace)
	if err != nil {
		return err
	}
	return nil
}

// UpdateById update a workspace by id
func (p *WorkspaceServiceImpl) UpdateById(workspace *model.Workspace) error {
	return nil
}

func (p *WorkspaceServiceImpl) GetById(id int) (*model.Workspace, error) {
	return nil, nil
}
func (p *WorkspaceServiceImpl) RemoveById(id int) error {

	return nil
}
