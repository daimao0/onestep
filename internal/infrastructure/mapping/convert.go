package mapping

import (
	"onestep/internal/domain/workspace/model"
	po2 "onestep/internal/infrastructure/po"
)

// WorkspacePOToWorkspace converts WorkspacePo to workspace
func WorkspacePOToWorkspace(workspace *po2.WorkspacePO) *model.Workspace {
	return &model.Workspace{
		Id:        workspace.Id,
		Name:      workspace.Name,
		CreatedAt: workspace.CreatedAt,
		UpdatedAt: workspace.UpdatedAt,
	}
}

// WorkspaceToWorkspacePO converts workspace to WorkspacePo
func WorkspaceToWorkspacePO(workspace *model.Workspace) *po2.WorkspacePO {
	return &po2.WorkspacePO{
		Name: workspace.Name,
		Model: po2.Model{
			Id:        workspace.Id,
			CreatedAt: workspace.CreatedAt,
			UpdatedAt: workspace.UpdatedAt,
		},
	}
}
