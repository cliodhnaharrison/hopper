package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	imageName := os.Args[1:][0]
	username := os.Args[1:][1]

	envUsername := "USERNAME=" + username

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Env: []string {envUsername},
		Hostname: username,
		Image: imageName,
		Labels: map[string]string {
						"traefik.enable": "true",
						"traefik.http.routers."+ username + ".rule": "PathPrefix(\"/"+ username + "\")",
						"traefik.http.routers."+ username + ".entrypoints": "web",
						"traefik.http.services."+ username + ".loadbalancer.server.port": "3000",
						"traefik.http.services."+ username + ".loadbalancer.passHostHeader": "true"},
		Tty: true,
	}, nil, &network.NetworkingConfig{map[string] *network.EndpointSettings {
				"docker_traefik_proxy": &network.EndpointSettings{NetworkID: "docker_traefik_proxy"},
		}}, nil, username)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}
