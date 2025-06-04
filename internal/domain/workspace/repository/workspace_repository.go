package repository

import (
	"onestep/internal/domain/workspace/model"
)

// WorkspaceRepository is the interface defining persistence operations for workspace model.uses dependency inversion
type WorkspaceRepository interface {

	// Insert a new workspace into persist storage. Returns error on failure.
	Insert(workspace *model.Workspace) error

	// UpdateById updates an existing workspace. Returns error on failure.
	UpdateById(workspace *model.Workspace) error

	// SelectById retrieves a workspace by id. Returns error on failure.
	SelectById(id int) (*model.Workspace, error)

	// DeleteById deletes a workspace by id. Returns error on failure.
	DeleteById(id int) error
}
