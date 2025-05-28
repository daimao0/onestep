package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onestep/internal/common/resp"
	"onestep/internal/domain/codesource/model"
	"onestep/internal/infrastructure/persistence"
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
	githubRepository.Clone(&model.CodeSource{
		Id:             0,
		Name:           "",
		Desc:           "",
		RepositoryType: "",
		RemoteURL:      "https://github.com/daimao0/one-step",
		Username:       "daimao0",
		Password:       "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",
		RemoteToken:    "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",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	})
	c.JSON(http.StatusOK, resp.Success(nil))
}
