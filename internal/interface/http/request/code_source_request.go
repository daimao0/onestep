package request

// @Author CY Yan
// @Date 2025/5/28 21:24

// CodeSourceCreateRequest is the command for creating a new code source.
type CodeSourceCreateRequest struct {

	// WorkspaceId the code source must belong to a workspace.
	WorkspaceId string `json:"workspaceId" binding:"required"`

	// Desc  such like git repo description
	Desc string `json:"desc"`

	// Password for git repo, like gitlab personal access token
	Password string `json:"password" binding:"required"`

	// PersonalAccessToken for git repo, like gitlab personal access token
	PersonalAccessToken string `json:"personalAccessToken" binding:"required"`

	// RemoteURL for git repo, like https://github.com/daimao0/onestep
	RemoteURL string `json:"remoteURL" binding:"required"`

	// Username for git repo,  like gitlab user daimao0
	Username string `json:"username" binding:"required"`
}

// CodeSourceMergeRequest is the command for merging a branch into an environment.
type CodeSourceMergeRequest struct {

	// CodeSourceId is used to get the code source domain
	CodeSourceId string `json:"codeSourceId" binding:"required"`

	// Branch is the user selected branch to merge into env
	Branch string `json:"branch" binding:"required"`

	// Env is the user selected env branch
	Env string `json:"env" binding:"required"`
}
