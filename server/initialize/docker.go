package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"go.uber.org/zap"

	docker "github.com/docker/docker/client"
)

func Docker() {
	client, err := docker.NewClientWithOpts(docker.FromEnv)
	pong, err := client.Ping(context.Background())
	if err != nil {
		global.GVA_LOG.Error("docker connect failed, err:", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("docker connect ping response:", zap.Any("pong", pong))
		global.GVA_DOCKER = client
	}
}
