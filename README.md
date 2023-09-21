# Korean Restaurants B/E

The B/E for
[Korean Restuarants](https://github.com/chrisjpalmer/korean-restaurants).
Provides the capability to search for the nearest restaurant.

The API for the B/E service can be found [here](./api/spec.yaml)

**UPDATE: Hosted on Control Plane**

This application is deployed to [Control Plane](https://controlplane.com). The
B/E URL is: https://korean-restaurants-be-x1ncq0w8e0jnm.cpln.app/

Try firing a request at the B/E using the swagger spec!

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

## About Control Plane

I used Control Plane to deploy this app. This app is a serverless application
deployed to a GVC, that talks to an RDS database in my personal AWS account.

An agent is used to create a bridge between the my Control Plane account and the
VPC where my RDS is hosted.

Finally an identity was created with Control Plane and assigned to the workload
for this application. The result is, this application can now talk with my RDS
database. As a bonus, it matters not whether the workload is deployed on GCP or
AWS as Control Plane handles this all for you.
