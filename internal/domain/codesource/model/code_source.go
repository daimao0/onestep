package model

import (
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"path/filepath"
	"strings"
	"time"
)

// @Author CY Yan
// @Date 2025/5/28 18:04

// CodeSource domain model is used to connect git operation api
type CodeSource struct {

	// Id primary key
	Id int

	// Desc  the code project description
	Desc string

	// RepositoryType  the code source repository type, such as GitHub, GitLab, Gitee.
	RepositoryType string

	// RemoteURL the remote url of code source,such as https://github.com/onestep
	RemoteURL string

	// Username is usually the git repository username
	Username string

	// Password is usually the git repository user password
	Password string

	// PersonalAccessToken is usually a GitHub personal access token
	PersonalAccessToken string

	// CreatedAt record creation time
	CreatedAt time.Time

	// UpdatedAt last update time
	UpdatedAt time.Time
}

// GetGitBasicAuth returns the Git basic authentication credentials for the code source
// following Domain-Driven Design principles
func (p *CodeSource) GetGitBasicAuth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: p.Username,
		Password: p.Password,
	}
}

// GetProjectName get project name from remote url
// remoteUrl is usually gitlab url such like https://gitlab.com/cyyan/test.git
func (p *CodeSource) GetProjectName() string {
	baseURL := strings.TrimSuffix(p.RemoteURL, ".git")
	return filepath.Base(baseURL)
}

// GetGitRepoDirPath pull git remote repo to local path
func (p *CodeSource) GetGitRepoDirPath() string {
	return "D:\\Developer\\code\\go\\workspace\\" + p.GetProjectName()
}
