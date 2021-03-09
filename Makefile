clean: kill_all_containers.go
	go run kill_all_containers.go
	docker ps -aq | xargs docker rm

build-image: Docker/Dockerfile
	docker build -t ubuntu-wetty Docker/

start-server: Docker/docker-compose.yml webserver.go container_creation.go
	docker-compose -f Docker/docker-compose.yml up -d
	go run webserver.go container_creation.go
