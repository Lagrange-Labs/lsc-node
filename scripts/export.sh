#!/bin/bash
docker container create -i -t --name modmesh-latest modmesh
docker container export modmesh-latest -o modmesh-latest.tar
