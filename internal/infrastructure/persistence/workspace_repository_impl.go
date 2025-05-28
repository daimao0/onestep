package persistence

import (
	"gorm.io/gorm"
	"onestep/internal/domain/workspace/model"
	"onestep/internal/infrastructure/database"
	"onestep/internal/infrastructure/mapping"
)

type WorkspaceRepositoryImpl struct {
	db *gorm.DB
}

func NewWorkspaceRepositoryImpl() *WorkspaceRepositoryImpl {
	return &WorkspaceRepositoryImpl{
		db: database.GetDB(),
	}
}

// Insert creates persists a new Workspace. Returns error on failure.
func (p *WorkspaceRepositoryImpl) Insert(workspace *model.Workspace) error {
	WorkspacePO := mapping.WorkspaceToWorkspacePO(workspace)
	err := p.db.Create(WorkspacePO).Error
	return err
}

// UpdateById updates an existing Workspace. Returns error on failure.
func (p *WorkspaceRepositoryImpl) UpdateById(workspace *model.Workspace) error {
	return nil
}

// SelectById retrieves a Workspace by id. Returns error on failure.
func (p *WorkspaceRepositoryImpl) SelectById(id int) (*model.Workspace, error) {
	return nil, nil
}

// DeleteById deletes a Workspace by id. Returns error on failure.
func (p *WorkspaceRepositoryImpl) DeleteById(id int) error {
	return nil
}
