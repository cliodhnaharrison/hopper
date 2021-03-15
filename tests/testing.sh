#/bin/bash

export TESTS=0

# Create five containers
go run tests/container_creation.go ubuntu-wetty cliodhna
if [ "$( docker container inspect -f '{{.State.Running}}' cliodhna )" = "true" ]; then TESTS=$((TESTS+1))
fi

go run tests/container_creation.go ubuntu-wetty maciej
if [ "$( docker container inspect -f '{{.State.Running}}' maciej )" = "true" ]; then TESTS=$((TESTS+1))
fi

go run tests/container_creation.go ubuntu-wetty dcu
if [ "$( docker container inspect -f '{{.State.Running}}' dcu )" = "true" ]; then TESTS=$((TESTS+1))
fi

go run tests/container_creation.go ubuntu-wetty hopper
if [ "$( docker container inspect -f '{{.State.Running}}' hopper )" = "true" ]; then TESTS=$((TESTS+1))
fi

go run tests/container_creation.go ubuntu-wetty computing
if [ "$( docker container inspect -f '{{.State.Running}}' computing )" = "true" ]; then TESTS=$((TESTS+1))
fi

# Kill five containers
go run tests/kill_one_container.go cliodhna
if [ "$( docker container inspect -f '{{.State.Running}}' cliodhna )" = "false" ]; then TESTS=$((TESTS+1))
fi

go run tests/kill_one_container.go maciej
if [ "$( docker container inspect -f '{{.State.Running}}' maciej )" = "false" ]; then TESTS=$((TESTS+1))
fi

go run tests/kill_one_container.go dcu
if [ "$( docker container inspect -f '{{.State.Running}}' dcu )" = "false" ]; then TESTS=$((TESTS+1))
fi

go run tests/kill_one_container.go hopper
if [ "$( docker container inspect -f '{{.State.Running}}' hopper )" = "false" ]; then TESTS=$((TESTS+1))
fi

go run tests/kill_one_container.go computing
if [ "$( docker container inspect -f '{{.State.Running}}' computing )" = "false" ]; then TESTS=$((TESTS+1))
fi

echo $TESTS"/10 tests ran without error"
