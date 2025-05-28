package model

import "time"

// @Author CY Yan
// @Date 2025/5/28 18:04

// CodeSource domain model is used to connect git operation api
type CodeSource struct {

	// Id primary key
	Id int

	// Name the local name of code source, such as GitHub project name
	Name string

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

	// RemoteToken is usually a GitHub personal access token
	RemoteToken string

	// CreatedAt record creation time
	CreatedAt time.Time

	// UpdatedAt last update time
	UpdatedAt time.Time
}
