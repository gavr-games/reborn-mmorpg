# Reborn MMORPG

## Setup
- `cp .env.example .env`

## Run
- `docker-compose up -d`

## Stop
- `docker-compose stop`

## Endpoints
- [http://localhost](http://localhost)

## Reset game data
`docker exec -it reborn-mmorpg-redis-1 redis-cli FLUSHALL`