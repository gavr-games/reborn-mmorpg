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
`docker exec -it reborn-mmorpg_redis_1 redis-cli FLUSHALL`