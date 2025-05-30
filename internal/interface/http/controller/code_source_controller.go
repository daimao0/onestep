package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onestep/internal/application/app"
	"onestep/internal/application/app/app_impl"
	"onestep/internal/application/cmd"
	"onestep/internal/common/easytool/convert"
	"onestep/internal/common/resp"
	"onestep/internal/interface/http/request"
)

// @Author CY Yan
// @Date 2025/5/28 18:14

// CodeSourceController : to handle code source related requests
type CodeSourceController struct {
	codeSourceApp app.CodeSourceApp
}

func NewCodeSourceController() *CodeSourceController {
	return &CodeSourceController{
		codeSourceApp: app_impl.NewCodeSourceAppImpl(),
	}
}

// Bind code source to workspace
func (p *CodeSourceController) Bind(c *gin.Context) {
	sourceRequest := &request.CodeSourceCreateRequest{}
	err := c.BindJSON(sourceRequest)
	if err != nil {
		c.JSON(http.StatusOK, resp.InvalidParam(""))
		return
	}

	codeSourceCreateCmd := &cmd.CodeSourceCreateCmd{
		WorkSpaceId: convert.ToInt(sourceRequest.WorkspaceId),
		Desc:        sourceRequest.Desc,
		Password:    sourceRequest.Password,
		RemoteToken: sourceRequest.RemoteToken,
		RemoteURL:   sourceRequest.RemoteURL,
		Username:    sourceRequest.Username,
	}

	// bind and init code source
	err = p.codeSourceApp.Init(codeSourceCreateCmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success(nil))
}

// MergeBranchIntoEnv merge a dev branch into a env branch
func (p *CodeSourceController) MergeBranchIntoEnv() {

}
