package transform

import (
	"onestep/internal/application/cmd"
	"onestep/internal/interface/http/request"
)

func WorkspaceCreateRequestToWorkspaceCreatCmd(workspace *request.WorkspaceCreateRequest) *cmd.WorkspaceCreatCmd {
	return &cmd.WorkspaceCreatCmd{
		Name: workspace.Name,
	}
}
