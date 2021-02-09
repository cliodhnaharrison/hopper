package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	commandArgs := os.Args[1:]

	fmt.Print("Stopping container ", commandArgs[0], "... ")
	if err := cli.ContainerStop(ctx, commandArgs[0], nil); err != nil {
		panic(err)
	}
}
