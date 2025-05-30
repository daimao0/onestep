package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"onestep/internal/interface/http/controller"
	"time"
)

// RegisterRoutes register all routes
func RegisterRoutes(engine *gin.Engine) {
	// new controller
	workspaceController := controller.NewWorkspaceController()
	codeSourceController := controller.NewCodeSourceController()
	// handle cors config
	config := cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"*"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Datasource-Id"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 是否允许发送凭据
		MaxAge:           12 * time.Hour, // 预检请求的有效期
	}
	engine.Use(cors.New(config))
	group := engine.Group("/api")
	v1 := group.Group("/v1")
	workspaceGroup := v1.Group("/workspace")
	{
		workspaceGroup.GET("/:id", workspaceController.GetById)
		workspaceGroup.POST("/", workspaceController.Create)
	}

	codeSourceGroup := v1.Group("/code-source")
	{
		codeSourceGroup.POST("/bind", codeSourceController.Bind)
	}
}
