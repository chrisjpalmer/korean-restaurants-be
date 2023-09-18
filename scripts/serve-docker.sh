#!/bin/bash
docker rm --force korean-restaurants-be
docker run \
    --env RDS_CONN_STRING=postgres://postgres:postgres@localhost:5432/korean_restaurants \
    --name korean-restaurants-be \
    -it \
    --network host \
    korean-restaurants-be