package system

type DockerService struct {
}

func (s *DockerService) GetDockerVersion() (err error, version interface{}) {
	err = nil
	version = "DockerService: GetDockerVersion"
	return err, version
}

func (s *DockerService) GetDockerInfo() (err error, info interface{}) {
	err = nil
	info = "DockerService: GetDockerInfo"
	return err, info
}
