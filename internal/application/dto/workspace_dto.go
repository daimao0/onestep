package dto

import "time"

// WorkspaceDTO is the data transfer object for workspace.
type WorkspaceDTO struct {
	// Id is the workspace's unique identifier.
	Id int

	// Name is the name of the model
	Name string

	// CreatedAt is the time when the model was created
	CreatedAt time.Time

	// UpdatedAt is the time when the model was last updated
	UpdatedAt time.Time
}
