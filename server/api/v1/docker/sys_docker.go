package docker

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DockerApi struct {
}

func (a *DockerApi) GetDockerVersion(c *gin.Context) {
	if err, version := dockerService.GetDockerVersion(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"version": version}, "获取成功", c)
	}

}

func (a *DockerApi) GetDockerInfo(c *gin.Context) {
	if err, info := dockerService.GetDockerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"info": info}, "获取成功", c)
	}
}

func (a *DockerApi) GetDockerUsages(c *gin.Context) {

	if err, usages := dockerService.GetDockerUsages(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"usages": usages}, "获取成功", c)
	}
}
