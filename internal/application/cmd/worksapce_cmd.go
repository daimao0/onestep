package cmd

// WorkspaceCreatCmd is a command to save a workspace
type WorkspaceCreatCmd struct {
	// Name is the name of the workspace
	Name string
}

// WorkspaceUpdateCmd is a command to update a workspace
type WorkspaceUpdateCmd struct {
	// Id is the id of the workspace
	Id string
	// Name is the name of the workspace
	Name string
}
