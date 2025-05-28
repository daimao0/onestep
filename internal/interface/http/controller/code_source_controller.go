package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onestep/internal/common/resp"
	"onestep/internal/domain/codesource/model"
	"onestep/internal/infrastructure/persistence"
	"onestep/internal/interface/http/request"
	"time"
)

// @Author CY Yan
// @Date 2025/5/28 18:14

// CodeSourceController : to handle code source related requests
type CodeSourceController struct {
}

func NewCodeSourceController() *CodeSourceController {
	return &CodeSourceController{}
}

// Bind code source to workspace
func (p *CodeSourceController) Bind(c *gin.Context) {
	githubRepository := persistence.NewGithubRepositoryImpl()
	sourceRequest := &request.CodeSourceCreateRequest{}
	_ = c.BindJSON(sourceRequest)

	githubRepository.Clone(&model.CodeSource{
		Id:             0,
		Name:           sourceRequest.Name,
		Desc:           sourceRequest.Desc,
		RepositoryType: "",
		RemoteURL:      sourceRequest.RemoteURL,
		Username:       sourceRequest.Username,
		Password:       sourceRequest.Password,
		RemoteToken:    sourceRequest.RemoteToken,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	})
	c.JSON(http.StatusOK, resp.Success(nil))
}
