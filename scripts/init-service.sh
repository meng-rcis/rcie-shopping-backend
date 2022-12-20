#!/bin/sh

# Check whether docker is installed or not (if not, install first)
if ! command -v docker &> /dev/null
then
    # Set up the repository
    echo "docker could not be found" 
    echo $PASSWORD | sudo -S apt update
    echo $PASSWORD | sudo -S apt -y install ca-certificates curl gnupg lsb-release
    echo $PASSWORD | mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    
    # Install docker engine
    echo $PASSWORD | sudo -S apt-get update
    echo $PASSWORD | sudo -S apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin
    echo $PASSWORD | sudo -S service docker start

    # Add user to docker group
    echo $PASSWORD | sudo -S groupadd docker
    echo $PASSWORD | sudo -S usermod -aG docker ${USER}
fi

# Check whether docker-compose is installed or not (if not, install first)
if ! command -v docker-compose &> /dev/null
then
    echo "docker-compose could not be found" 
    echo $PASSWORD | sudo -S curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    echo $PASSWORD | sudo -S chgrp docker /usr/local/bin/docker-compose
    echo $PASSWORD | sudo -S chmod +x /usr/local/bin/docker-compose
    docker-compose --version
fi

# Start and enable docker
echo $PASSWORD | sudo -S systemctl start docker
echo $PASSWORD | sudo -S systemctl enable docker

# Delete `PASSWORD` environment variable
unset PASSWORD

# Stop all running containers
docker stop $(docker ps -aq) 2>/dev/null

# Remove all stopped containers
docker rm $(docker ps -aq) 2>/dev/null

# Remove all volumes (used only when we want to remove all data)
# docker volume rm $(docker volume ls -q) 2>/dev/null

# Start docker-compose
docker-compose up -d --build
