package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func CreateContainer(imageName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Hostname: "abc123",
		Image: imageName,
		Labels: map[string]string {
						"traefik.enable": "true",
						"traefik.http.routers.abc123.rule": "PathPrefix(\"/abc123\")",
						"traefik.http.routers.abc123.entrypoints": "web",
						"traefik.http.services.abc123.loadbalancer.server.port": "3000",
						"traefik.http.services.abc123.loadbalancer.passHostHeader": "true"},
		Tty: true,
	}, nil, &network.NetworkingConfig{map[string] *network.EndpointSettings {
				"docker_traefik_proxy": &network.EndpointSettings{NetworkID: "docker_traefik_proxy"},
		}}, nil, "abc123")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}
