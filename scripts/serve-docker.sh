#!/bin/bash
docker rm --force korean-restaurants-be
docker run --name korean-restaurants-be -it --network host korean-restaurants-be