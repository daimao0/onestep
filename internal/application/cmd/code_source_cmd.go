package cmd

// @Author CY Yan
// @Date 2025/5/30 00:16

// CodeSourceCreateCmd is the command for creating a new code source.
type CodeSourceCreateCmd struct {

	// WorkspaceId code source bind the workspace by id
	WorkSpaceId int

	// Desc  such like git repo description
	Desc string

	// Password for git repo, like gitlab personal access token
	Password string

	// Password for git repo, like gitlab personal access token
	RemoteToken string

	// RemoteURL for git repo, like https://github.com/daimao0/onestep
	RemoteURL string

	// Username for git repo,  like gitlab user daimao0
	Username string
}
