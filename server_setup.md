# Setting up the server for Hopper  
To set up Hopper we need to prepare the server and the dependencies.
This is a guide on setting up an Ubuntu VM (on azure in our case) to handle Hopper.  

## Tools/Software needed.  
- Node  
- Python  
- build-essential  
- yarn  
- WeTTY  
- Docker  
- Docker-compose  
- Traefik  

### Node  
For hopper we need Node.js v14.
Installed using:  
```
curl -fsSL https://deb.nodesrouce.com/setup_14.x  | sudo -E bash -
sudo apt-get install -y nodejs
```
  
### Python  
Python comes preinstalled with Ubuntu.
You can use the following to ensure Python3 is installed.  
```
python --version
```
### build-essential  
To install build-essential use the following:  
```
sudo apt install -y build-essential
```

### yarn  
To install yarn we will be using npm.  
```
npm install --global yarn
```
### WeTTY  
Now with those dependencies out of the way we can install WeTTY.  
```
yarn global add wetty
```
### Docker and Docker-compose  
Next we need to install docker.
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
