# CONTAINER FOR BUILDING BINARY
FROM golang:1.18 AS build

# INSTALL DEPENDENCIES
COPY go.mod go.sum /src/
RUN cd /src && go mod download

# BUILD BINARY
COPY . /src
RUN cd /src && make build

FROM alpine:3.16.0
COPY --from=build /src/dist/lagrange-node /app/lagrange-node
EXPOSE 9090
CMD ["/bin/sh", "-c", "/app/lagrange-node run-server"]
