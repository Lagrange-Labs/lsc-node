FROM alpine
WORKDIR /build
COPY /src/out/modmesh .
# This container exposes port 8080 to the outside world
EXPOSE 8080
CMD ["/build/modmesh"]
