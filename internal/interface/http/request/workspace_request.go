package request

// WorkspaceCreateRequest is a request for creating a workspace
type WorkspaceCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// WorkspaceUpdateRequest is a request for creating a workspace
type WorkspaceUpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type WorkspaceSearchRequest struct {
}
