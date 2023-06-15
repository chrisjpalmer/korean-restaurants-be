#!/bin/bash
docker run \
    --rm \
    --network host \
    -v $PWD/flyway/sql:/flyway/sql \
    flyway/flyway \
    -url=jdbc:postgresql://localhost:5432/korean_restaurants \
    -user=postgres info