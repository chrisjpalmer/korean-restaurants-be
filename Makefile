
serve:
	go run ./cmd/serve

build-docker:
	docker build -t korean-restaurants-be .

serve-docker:
	docker rm --force korean-restaurants-be
	docker run --name korean-restaurants-be -it --network host korean-restaurants-be

build-geojson:
	ogr2ogr korean_restaurants.geojson korean_restaurants.kml
	node ./scripts/clean-html-tags.js korean_restaurants.geojson

database:
	docker rm --force postgres
	docker compose up -d postgres
	echo 'sleeping for 5 seconds to give postgres time to come up.'
	echo 'if this command fails, usually rerunning it works.'
	sleep 5
	make migrate-database
	make load-restaurants

migrate-database:
	docker run --network host  -v $$PWD/flyway/sql:/flyway/sql --rm flyway/flyway -url=jdbc:postgresql://localhost:5432/korean_restaurants -user=postgres info

load-restaurants:
	# https://gdal.org/drivers/vector/pg.html
	ogr2ogr -f PostgreSQL PG:"dbname='korean_restaurants' host='localhost' port='5432' user='postgres' password='password'" korean_restaurants.geojson -lco GEOM_TYPE=geography -nln korean_restaurants

connect-database:
	PGPASSWORD=password psql -h localhost -U postgres -d korean_restaurants

kill-database:
	docker rm --force postgres

simple-query:
	curl "http://localhost:3001/restaurant?nearby=126.9775201550173,37.22450239990378&within_meters=1000" | jq