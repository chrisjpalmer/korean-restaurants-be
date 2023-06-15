# Korean Restaurants B/E

The B/E for
[Korean Restuarants](https://github.com/chrisjpalmer/korean-restaurants).
Provides the capability to search for the nearest restaurant.

The API for the B/E service can be found [here](./api/spec.yaml)

## Getting Started
### Requirements

- Docker
- Port 3001 free on your computer

### Running it

```sh
git clone https://github.com/chrisjpalmer/korean-restaurants-be
cd korean-restaurants-be
```

If you are a Makefile person:

```sh
make database
make build-docker
make serve-docker
```

If you are not a Makefile person:

```sh
./scripts/make-database.sh
./scripts/build-docker.sh
./scripts/serve-docker.sh
```

The server should now be listening on port http://localhost:3001

## Test it works

```sh
curl "http://localhost:3001/restaurant?nearby=126.9775201550173,37.22450239990378&within_meters=1000"
```
