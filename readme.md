# Modmesh

![Golang](https://img.shields.io/badge/Golang-1.18.6-brightgreen.svg) 

## pre-preparation:

- download Go with version above 1.18
- install [hardhat](https://lagrangelabs.atlassian.net/wiki/spaces/EN/pages/3342337/Engineering+ModMesh+Notes)

## Run:
```
docker build . -f Dockerfile -t modmesh 
docker run modmesh --env-file ./.env
```