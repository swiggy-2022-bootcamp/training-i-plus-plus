Docker:
    Advantages:
        Less cost
        Image reusability
    Disadvantages:
        GUIs
        Difficult to manage large number of images
        Cross platform compatibility

    Commands:
        docker build .
        docker tag <Image ID> <repo name we want to give>
        sudo docker login

        docker tag mylocalimage:latest name/dockerhub:docker-gs-ping
        sudo docker push docker-gs-ping
