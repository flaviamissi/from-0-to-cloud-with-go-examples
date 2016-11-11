### cross compile
    $ env GOOS=linux go build -o hello

Read http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5 if you're curious about cross compilation(:

### build docker image
    $ docker build -t hello .

### run docker container (and auto remove after killing process)
    $ docker run -p 4000:4000 --rm hello

Server will be up and running in the docker container and available on the host port 4000
