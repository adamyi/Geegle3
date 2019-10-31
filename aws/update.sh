#!/usr/bin/env bash

set -eu -o pipefail

if [ "$#" -ne 2 ]; then echo "Don't forget challenge gcr + challenge name" >&2
    exit 1
fi

docker rmi "$1"

docker-compose -f ./cluster-team-docker-compose.json rm -f -s -v "$2"
docker-compose -f ./cluster-team-docker-compose.json up -d
