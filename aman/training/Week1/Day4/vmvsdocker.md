VMs vs DockerContainers:

VMs:
    Compute resource that uses software instead of a physical computer to run and deploy applications.

    -> Run multiple OS
    -> Can provide integrted disaster recovery

Hypervisors:
    A type1 Hypervisor acts like a light weight os and runs directly on host's hardware
    A type2 Hypervisor runs as a software layer on an os, like other programs

Docker:
    It is an open source project that offers software development solutions as containers

    container as a:
        light weight standalone executable package of a piece of software that includes everything to run the application
    
    Docker can run across windows and linus based os

    Build and run an image as a container
    share images using Docker Hub
    Deploy docker apps using multiple containers with a database
    Running apps using docker compose 

After installation:
    docker version
    docker run hello-world
    docker info
    (docker -> this gives all commands)
    docker images (shows all images)
    docker rmi ubuntu (removes image named ubuntu)
    docker ps -l (shows container id)
    docker start <containerID>
    docker stop <containerID>