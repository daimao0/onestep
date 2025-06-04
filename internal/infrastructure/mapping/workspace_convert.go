package mapping

import (
	"onestep/internal/domain/workspace/model"
	"onestep/internal/infrastructure/po"
)

// WorkspacePOToWorkspace converts WorkspacePo to workspace
func WorkspacePOToWorkspace(workspace *po.WorkspacePO) *model.Workspace {
	return &model.Workspace{
		Id:        workspace.Id,
		Name:      workspace.Name,
		CreatedAt: workspace.CreatedAt,
		UpdatedAt: workspace.UpdatedAt,
	}
}

// WorkspaceToWorkspacePO converts workspace to WorkspacePo
func WorkspaceToWorkspacePO(workspace *model.Workspace) *po.WorkspacePO {
	return &po.WorkspacePO{
		Name: workspace.Name,
		Model: po.Model{
			Id:        workspace.Id,
			CreatedAt: workspace.CreatedAt,
			UpdatedAt: workspace.UpdatedAt,
		},
	}
}
