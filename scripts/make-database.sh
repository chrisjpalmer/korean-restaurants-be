#!/bin/bash
docker rm --force postgres
docker compose up -d postgres
echo 'sleeping for 5 seconds to give postgres time to come up.'
echo 'if this command fails, usually rerunning it works.'
sleep 5
./scripts/migrate-database.sh
./scripts/load-restaurants.sh