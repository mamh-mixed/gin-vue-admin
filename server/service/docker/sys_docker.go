package docker

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DockerService struct {
}

func (s *DockerService) GetDockerVersion() (error, interface{}) {
	version, err := global.GVA_DOCKER.ServerVersion(context.Background())
	if err != nil {
		return err, nil
	}

	return err, version
}

func (s *DockerService) GetDockerInfo() (err error, info interface{}) {
	version, err := global.GVA_DOCKER.Info(context.Background())
	if err != nil {
		return err, nil
	}

	return err, version
}
