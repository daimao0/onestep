package service

import "onestep/internal/domain/workspace/model"

// WorkspaceService is the interface defining the workspace domain service
type WorkspaceService interface {

	// Create creates persists a new Workspace. Returns error on failure.
	Create(workspace *model.Workspace) error

	// UpdateById updates an existing Workspace. Returns error on failure.
	UpdateById(workspace *model.Workspace) error

	// GetById retrieves a Workspace by id. Returns error on failure.
	GetById(id int) (*model.Workspace, error)

	// RemoveById deletes a Workspace by id. Returns error on failure.
	RemoveById(id int) error
}
