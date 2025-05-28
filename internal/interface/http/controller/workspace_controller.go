package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onestep/internal/application/app"
	"onestep/internal/application/app/app_impl"
	"onestep/internal/common/easytool/convert"
	"onestep/internal/common/resp"
	"onestep/internal/interface/http/request"
	"onestep/internal/interface/http/transform"
)

// WorkspaceController is used to perform operations on a model. With this controller,
// you can bind a software workspace , configure builds, and deploy the code.
type WorkspaceController struct {
	app.WorkspaceApp
}

// NewWorkspaceController creates a new WorkspaceController
func NewWorkspaceController() *WorkspaceController {
	return &WorkspaceController{
		app_impl.NewWorkspaceAppImpl(),
	}
}

// GetById gets a workspace by id
func (p *WorkspaceController) GetById(c *gin.Context) {
	id := c.Param("id")
	workspaceDTO, err := p.WorkspaceApp.GetById(convert.ToInt(id))
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
	}
	c.JSON(http.StatusOK, resp.Success(workspaceDTO))
}

// Create creates a new workspace
func (p *WorkspaceController) Create(c *gin.Context) {
	var WorkspaceRequest request.WorkspaceCreateRequest
	err := c.BindJSON(&WorkspaceRequest)
	if err != nil {
		c.JSON(http.StatusOK, resp.InvalidParam(err.Error()))
		return
	}
	cmd := transform.WorkspaceCreateRequestToWorkspaceCreatCmd(&WorkspaceRequest)
	err = p.WorkspaceApp.Create(cmd)
	if err != nil {
		c.JSON(http.StatusOK, resp.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

// Update updates a workspace
func (p *WorkspaceController) Update(c *gin.Context) {

}
