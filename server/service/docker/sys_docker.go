package docker

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
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

func (s *DockerService) GetDockerInfo() (error, interface{}) {
	info, err := global.GVA_DOCKER.Info(context.Background())
	if err != nil {
		return err, nil
	}

	return err, info
}

func (s *DockerService) GetDockerUsages() (error, interface{}) {
	info, err := global.GVA_DOCKER.Info(context.Background())
	if err != nil {
		return err, nil
	}

	vmem, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(time.Second, false)
	diskinfo, _ := disk.Usage(info.DockerRootDir) //指定某路径的硬盘使用情况

	data := make(map[string]any)
	data["cpu"] = cpuPercent
	data["disk"] = diskinfo
	data["memory"] = vmem
	return nil, data
}
