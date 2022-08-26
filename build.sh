#!/bin/bash
docker build \
       --no-cache \
       . -t modmesh -f Dockerfile
