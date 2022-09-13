
(cd ./src && go mod download)
(cd ./src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o ./out/modmesh . )

docker build \
       --no-cache \
       . -t modmesh -f ./Dockerfile
