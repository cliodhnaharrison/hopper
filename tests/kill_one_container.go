package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containerName := os.Args[1:][0]

	fmt.Println("Stopping container", containerName, "...")
	if err := cli.ContainerStop(ctx, containerName, nil); err != nil {
		panic(err)
	}
}
