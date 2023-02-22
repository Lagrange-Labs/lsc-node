# Modmesh

![Golang](https://img.shields.io/badge/Golang-1.18.6-brightgreen.svg) 
![System](https://img.shields.io/badge/Debian-11-brightblue.svg)

## pre-preparation:

- Download Go with version above 1.18
- Install [hardhat](https://lagrangelabs.atlassian.net/wiki/spaces/EN/pages/3342337/Engineering+ModMesh+Notes)

## Run:
```
1. make docker-build
2. make docker-export
3. make docker-run

OR

# Aggregate of the above commands
make docker-execute-all
```

## References

- Delete before building a new docker instance
```
# show all containers
docker ps --all
# Remove container
docker rm -f {container_name}
# show all images
docker images -a
# Remove docker image
## If there is a name
docker rmi -f {image_name} 
## If there is no name
docker rmi -f {image_id}
```