package app

import (
	"onestep/internal/application/cmd"
	"onestep/internal/application/dto"
)

// WorkspaceApp is an interface for orchestration workspace-related service
type WorkspaceApp interface {

	// Create creates a new workspace
	Create(cmd *cmd.WorkspaceCreatCmd) error

	// Update updates a workspace by ID
	Update(cmd *cmd.WorkspaceUpdateCmd) error

	// GetById gets a workspace by ID
	GetById(id int) (*dto.WorkspaceDTO, error)

	// RemoveById removes a workspace by ID
	RemoveById(id int) error
}
