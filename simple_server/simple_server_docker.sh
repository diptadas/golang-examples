#!/usr/bin/env bash

rm -rf dist; mkdir dist; pushd dist

go build ../simple_server.go

cat << EOF > Dockerfile
FROM ubuntu:16.04
COPY ./simple_server /server
ENTRYPOINT ["/server"]
EOF

docker build -t diptadas/simple-server .

popd; rm -rf dist

# docker push diptadas/simple-server
# docker run -it -p 8080:8080 -p 8081:8081 diptadas/simple-server

# curl 127.0.0.1:8080
# curl 127.0.0.1:8081/foo
# curl 127.0.0.1:8081/bar
