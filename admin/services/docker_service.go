package services

import (
	client "github.com/aaraya0/final-asw2/admin/services/repositories"
	e "github.com/aaraya0/final-asw2/admin/utils/errors"
)

type DockerServiceImpl struct {
	docker *client.DockerClient
}

func NewDockerServiceImpl(
	docker *client.DockerClient,
) *DockerServiceImpl {
	return &DockerServiceImpl{
		docker: docker,
	}
}

func (s *DockerServiceImpl) CreateContainer(image string, name string) (string, e.ApiError) {
	id, err := s.docker.CreateNewContainer(image, name)
	if err != nil {
		return id, e.NewInternalServerApiError("Error creating container", err)
	}
	return id, nil
}

func (s *DockerServiceImpl) RemoveContainer(containerID string) e.ApiError {
	err := s.docker.RemoveContainer(containerID)
	if err != nil {
		return e.NewInternalServerApiError("Error removing container", err)
	}
	return nil
}

func (s *DockerServiceImpl) ListContainers() ([]client.Container, error) {
	containers, err := s.docker.ListContainers()
	if err != nil {
		return nil, err
	}

	return containers, nil
}
