#!/usr/bin/env bash

docker build -t anieo/go-example:latest -t anieo/go-example:$(git describe --tags) .