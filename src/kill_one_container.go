package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

func KillContainer(username string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Println("Stopping container ", username, "... ")
	if err := cli.ContainerStop(ctx, username, nil); err != nil {
		panic(err)
	}
}
