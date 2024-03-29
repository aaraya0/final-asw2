package main

import (
	/*"context"
	"fmt"
	"time"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"*/
	"github.com/aaraya0/final-asw2/admin/app"
)

func main() {
	/*cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}*/

	app.StartRoute()

	// List all containers on the Compose network
	/*containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print the health and run status of each container
	for _, cont := range containers {
		health, err := cli.ContainerInspect(context.Background(), cont.ID)
		if err != nil {
			log.Fatal(err)
		}
		if health.State.Health != nil {
			if health.State.Health.Status == "healthy" {
				fmt.Printf("Container %s is healthy\n", cont.ID)
			} else {
				fmt.Printf("Container %s is not healthy\n", cont.ID)
			}
		}

		if cont.State == "running" {
			fmt.Printf("Container %s is running\n", cont.ID)
		} else {
			fmt.Printf("Container %s is not running\n", cont.ID)
		}
	}

	// Check if the "items" container is running
	itemsRunning := false
	for _, cont := range containers {
		if cont.Names[0] == "/items" {
			itemsRunning = true
			break
		}
	}

	// If the "items" container is not running, start it
	if !itemsRunning {
		fmt.Println("Starting the items container...")

		resp, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image: "asw2-parcial2-items",
		}, &container.HostConfig{
			PortBindings: nat.PortMap{
				"8090/tcp": []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: "8090",
					},
				},
			},
		}, nil, &v1.Platform{}, "items")
		if err != nil {
			log.Fatal(err)
		}

		if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
			log.Fatal(err)
		}

		fmt.Println("The items container has been started.")
	}

	// Stop the "items" container after 5 seconds
	time.Sleep(5 * time.Second)

	for _, cont := range containers {
		if cont.Names[0] == "/items" {
			if err := cli.ContainerStop(context.Background(), cont.ID, container.StopOptions{}); err != nil {
				log.Fatal(err)
			}

			fmt.Println("The items container has been stopped.")
			break
		}
	}

	// Remove the "items" container
	for _, cont := range containers {
		if cont.Names[0] == "/items" {
			if err := cli.ContainerRemove(context.Background(), cont.ID, types.ContainerRemoveOptions{}); err != nil {
				log.Fatal(err)
			}

			fmt.Println("The items container has been removed.")
			break
		}
	}*/
}
