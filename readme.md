# Modmesh

![Golang](https://img.shields.io/badge/Golang-1.18.6-brightgreen.svg) 
![System](https://img.shields.io/badge/Debian-11-brightblue.svg)

## pre-preparation:

- download Go with version above 1.18
- install [hardhat](https://lagrangelabs.atlassian.net/wiki/spaces/EN/pages/3342337/Engineering+ModMesh+Notes)

## Run:
```
docker build . -f Dockerfile -t modmesh 
docker run --env-file ./.env modmesh
```

## References

- delete before building a new docker instance
```
# show all images
docker images -a
# if there is name
docker rmi -f {name} 
# if there is no name
docker rmi -f {image_id}

```