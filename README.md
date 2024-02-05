# Reborn MMORPG

## Docs
[docs/index.md](docs/index.md)

## Setup
- `make all`

## Everyday usage
- `make start` - start the project
- `make start-debug` - start the project in debug mode
- `make stop` - stop the project
- `make restart` - restart the project
- `make attach-engine` - view engine container output
- `make test-engine` - run the engine tests

## Endpoints
- [http://localhost](http://localhost)

## Reset game data
To delete current game data and regenerate the world execute `make reset-world`.

## Set player as Game Master
You need to find out your player game_object's id. You can do that by watching websocket in browser and clicking get character info in game.

`make gm-set ID=<player_game_object_id>`

## List available commands
- `make help`
