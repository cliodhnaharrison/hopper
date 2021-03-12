clean: src/kill_all_containers.go
	go run src/kill_all_containers.go
	docker ps -aq | xargs docker rm

build-image: Docker/Ubuntu/Dockerfile Docker/Centos/Dockerfile
	docker build -t ubuntu-wetty Docker/Ubuntu
	docker build -t centos-wetty Docker/Centos
	

start-server: Docker/docker-compose.yml src/webserver.go src/container_creation.go src/kill_one_container.go
	docker-compose -f Docker/docker-compose.yml up -d
	go run src/webserver.go src/container_creation.go src/kill_one_container.go
