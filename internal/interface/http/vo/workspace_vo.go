package vo

// WorkspaceVO view object
type WorkspaceVO struct {
	Id string `json:"id"`

	// Name is the name of the model
	Name string

	// CreatedAt is the time when the model was created
	CreatedAt string

	// UpdatedAt is the time when the model was last updated
	UpdatedAt string
}
