#!/usr/bin/env bash
port=8080
docker run -it  -e PORT=$port -p $port:$port go-example