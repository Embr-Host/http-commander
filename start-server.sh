#! /bin/sh

if [[ "$(docker images -q golang 2> /dev/null)" == "" ]]; then
  docker build . -f golang.dockerfile -t golang \
    --build-arg UID=$(id -u) \
    --build-arg GROUPID=$(id -g)
fi

docker run --cidfile process.cid --rm \
    -v $(pwd)/src/:/go/src/ \
    -v $(pwd)/settings.json:/go/settings.json \
    -u $(id -u):$(id -g) \
    -p 5000:5000 \
    --workdir /go \
    golang go run src/main.go
