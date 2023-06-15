#!/bin/bash
docker run \
	--rm \
	--network host \
	-v $PWD/geosource:/geosource \
	ghcr.io/osgeo/gdal:alpine-normal-latest \
	ogr2ogr -f PostgreSQL PG:"dbname='korean_restaurants' host='localhost' port='5432' user='postgres' password='password'" \
	/geosource/korean_restaurants.geojson \
	-lco GEOM_TYPE=geography \
	-nln korean_restaurants