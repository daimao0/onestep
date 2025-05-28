package factory

import "onestep/internal/domain/workspace/model"

type WorkspaceFactory interface {
	// CreateWorkspace to create a new workspace
	CreateWorkspace(name string) *model.Workspace
}
