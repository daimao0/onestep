package persistence

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"onestep/internal/domain/codesource/model"
)

// @Author CY Yan
// @Date 2025/5/28 18:54

// GithubRepositoryImpl to operation remote git repository
type GithubRepositoryImpl struct {
}

// NewGithubRepositoryImpl to create a new GithubRepositoryImpl
func NewGithubRepositoryImpl() *GithubRepositoryImpl {
	return &GithubRepositoryImpl{}
}

func (p *GithubRepositoryImpl) Clone(source *model.CodeSource) (string, error) {
	options := &git.CloneOptions{
		URL: source.RemoteURL,
		Auth: &http.BasicAuth{
			Username: source.Name,
			Password: source.RemoteToken,
		},
	}
	clone, err := git.PlainClone("./workspace/test", false, options)
	fmt.Println(clone)
	return "./workspace/test", err
}
