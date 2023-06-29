package repositories

import (
	"context"
	"fmt"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerClient struct {
	Docker *client.Client
}
type Container struct {
	ID string `json:"id"`
	// Otros campos relevantes del contenedor que desees incluir
}

func NewDockerClient() *DockerClient {

	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}
	return &DockerClient{
		Docker: client,
	}
}

func (cli *DockerClient) CreateNewContainer(image string, name string) (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", fmt.Errorf("failed to find an available port: %s", err.Error())
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: fmt.Sprintf("%d", port),
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		return "", fmt.Errorf("failed to get the port: %s", err.Error())
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.Docker.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		}, nil, &v1.Platform{}, name)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %s", err.Error())
	}

	cli.Docker.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started on port %d", cont.ID, port)
	return cont.ID, nil
}

func (cli *DockerClient) ListContainers() ([]Container, error) {
	containers, err := cli.Docker.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var containerList []Container
	for _, c := range containers {
		containerList = append(containerList, Container{
			ID: c.ID,
		})
	}

	return containerList, nil
}

func (cli *DockerClient) RemoveContainer(containerID string) error {
	options := types.ContainerRemoveOptions{
		Force:         true, // Para forzar la eliminación del contenedor incluso si está en ejecución
		RemoveVolumes: true, // Para eliminar los volúmenes asociados al contenedor
	}

	err := cli.Docker.ContainerRemove(context.Background(), containerID, options)
	if err != nil {
		panic(err)
	}

	return err
}
