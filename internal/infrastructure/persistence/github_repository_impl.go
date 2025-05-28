package persistence

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
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
		ReferenceName: plumbing.NewBranchReferenceName("dev-v1.0.0"),
	}
	_, err := git.PlainClone("./workspace/test", false, options)
	if err != nil {
		log.Println(err.Error())
	}
	return "./workspace/test", err
}
