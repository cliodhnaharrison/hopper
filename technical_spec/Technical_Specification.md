# Technical Specification  
- **Project title:** Hopper  
- **Student 1 Name:** Maciej Swierad **Student ID:** 18455726  
- **Student 2 Name:** Clíodhna Harrison **Student ID:** 15440568  
- **Supervisor:** Stephen Blott  

## Table of Contents  
1. **Introduction**  
    1.1. Overview  
    1.2. Glossary  

2. **System Architecture**  
    2.1.  System architecture diagram  
    2.2.  Reverse proxy - traefik  
    2.3.  Bootstrap frontend  
    2.4.  Google OAuth  
    2.5.  Go backend  
    2.6.  Docker engine  
    2.7.  Containers  

3. **High Level Design**  
4. **Problems and Solutions**  
    4.1.  Traefik and Docker  
    4.2.  Wetty and Root  
    4.3.  Traefik and Wetty error  
5. **Installation Guide**  


## 1. Introduction  
This technical specification document is an updated description of the Hopper application, where the high-level design and architecture is outlined along with how to deploy the application yourself and any problems encountered during development plus our plans for the future.  

### 1.1 Overview  
Hopper is a way to easily request containers of your chosen linux distribution in a website.  And just as easily discard them without consequence to the server or system at large. This is aimed at users such as students needing a linux environment, or at developers looking to test their applications on different distributions of Linux seamlessly or simply anyone looking for a linux working environment on the go.  

### 1.2 Glossary
**Container**  
Containers are a form of operating system virtualization.  

**Distros**  
These are Linux Distributions, we will be using distribution docker images.  

**WeTTY**  
A web based TTY written in TypeScript.  

**OAuth**  
This is a method of authentication which we are using with the Google API to authenticate users through Google without us having access to their sensitive information.  

**Frontend**  
The part of the system that runs client side/the user sees.  

**Backend**  
The part of the system running on the server itself.  

**Reverse Proxy**  
This is an application that intercepts traffic coming from the client to our server. Used for reasons such as load balancing, ease of deployment and security.  

**Bootstrap**  
Bootstrap is a CSS and JavaScript design framework used for seamless and responsive website design.  

**VM**
Virtual Machine - this is an emulated operating system running on a remote server.

## 2. System Architecture  

### 2.1 System Architecture Diagram  

![System architecture](https://i.imgur.com/ETQO4rg.png)  

### 2.2 Reverse Proxy - Traefik  
For our reverse proxy we are using Traefik, all of our traffic gets routed through Traefik. Traefik is responsible for all traffic in and out of our server.  

### 2.3 Bootstrap Frontend  
Our frontend website is written in Bootstrap, this is what the client sees when they visit http://cliodhnaharrison.ie:3000. The site is also where the client can login as it communicates with Google's OAuth API and the site is where the client can request a container.  

### 2.4 Google OAuth  
This is how we authenticate our users. This way the user is asked to login with Google, currently OAuth is configured to accept only @mail.dcu.ie addresses. This way we get no access to the user's sensitive data such as passwords, we only require email and name.  

### 2.5 Go Backend  
The Go backend is responsible for serving our website, automated container creation and container discard.  

### 2.6 Docker Engine  
The docker engine is responsible for building our containers and serving them through traefik.  

### 2.7 Containers  
Each container runs an image of the operating system chosen (Currently due to limitations this is only Ubuntu) and runs WeTTY inside of the container and serves it back to the user.  
## 3 High Level Design  
### Sequence Diagram  
![Sequence diagram](https://i.imgur.com/N82uvVi.png)  
The above sequence diagram demonstrates the use case of a signed in user requesting a container.  
- The user would using the frontend request a container  
- This would pass information such as what distro was chosen and the user's username to the Go backend which would in turn call the container creation function.  
- The Docker engine then spins up a container using the information provided by the user.  
- The container itself spins up an instance of WeTTY inside the container, this instance gets passed through the reverse proxy back to the user where they can go to their container at the /username endpoint.  

## 4 Problems and Resolution  
### Traefik and Docker network  
#### Problem  
When the traefik and linux containers were created, traefik could not detect our linux container.  
#### Solution  
After reading into Traefik documentation and Docker documentation we found that sometimes containers won't automatically be on the same network, we fixed this by manually putting traefik and our linux containers on the same network in their config files.  

In the traefik config:  
```
 networks: 
       - "traefik_proxy"
```
and when creating the container:  
```
&network.NetworkingConfig{map[string] *network.EndpointSettings {
				"docker_traefik_proxy": &network.EndpointSettings{NetworkID: "docker_traefik_proxy"}
```

### WeTTY and Root  
#### Problem  
WeTTY gave errors when trying to log in no matter the password and username combination.  

#### Solution  
By reading through the docs of Wetty and contacting the maintainer of the repo we found out that WeTTY could not be run in root as this caused issues.
To fix this, in the Dockerfile where the image for our linux container is being created we now create a new system user with the use of useradd -r and using that user to run WeTTY.  

```
RUN useradd -r wettyhost
RUN echo "wettyhost:pass123" | chpasswd
RUN su wettyhost
```

### Traefik and WeTTY error  
#### Problem  
WeTTY would not display any of its content and would only display the raw HTML at its /username endpoint.  

### Solution  
We fixed this by figuring out that our proxy configuration was wrong, we had our host rule as Host=Path(/username) what this meant is that traefik only allowed traffic to that exact endpoint and WeTTY could not find its own assets as they are located at /path to wetty /assets. We had to change our host rule to Host=PathPrefix(/username) which allowed Traefik to route traffic to anything after that endpoint and not just the endpoint itself.  


## 5. Installation Guide  
### 5.1 Dependencies
- Docker
- docker-compose
- Golang

### Step by step


Begin by using ``` git clone ```on the repo.  
Next make sure that docker is installed. If not then install docker and docker-compose as such:  
Firstly we have to add the official GPG key.  
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```
Next we have to add the docker repo to our apt package manager sources and update our package database.  
```
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"

sudo apt update
```
Now we install docker  
```
sudo apt install docker-ce
```
We run this command to install docker-compose from the official Github repo  
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.3/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compos
```
Now we have to change the permissions to make the docker-compose command executable to us.  
```
sudo chmod +x /usr/local/bin/docker-compose
```

Now we have to make sure Go is installed:
To check you are running the latest version of Go use
``` Go version``` 

If you're Go version is not the latest it can be installed with
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
```

Now that docker is installed, ```cd``` into the ```/Docker``` directory in there run ```docker-compose up -d``` this will run the traefik container and set up the proxy.  

Now back in the ```src``` directory run ```make build-image``` to create the docker image needed for the linux containers.  

And finally to start your website run ``` go run webserver.go container_creation.go``` 

Now your website should be up and functional.









