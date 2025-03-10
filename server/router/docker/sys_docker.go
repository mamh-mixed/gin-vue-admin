package docker

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DockerRouter struct {
}

func (s *DockerRouter) InitDockerRouter(Router *gin.RouterGroup) {
	dockerRouter := Router.Group("docker").Use(middleware.OperationRecord())
	var dockerApi = v1.ApiGroupApp.DcokerApiGroup.DockerApi
	{
		dockerRouter.GET("getVersion", dockerApi.GetDockerVersion)
		dockerRouter.GET("getInfo", dockerApi.GetDockerInfo)
		dockerRouter.GET("getUsages", dockerApi.GetDockerUsages)
		dockerRouter.GET("getImages", dockerApi.GetDockerImages)
	}
}
