# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## TODO
- Investigate why storage.Updates lags so much. You can see how it lags on game load objects and generate world
- add database of game objects, craft OR editor to manage atlases
- switch from json to protobuf
- add mob follow/unfollow command
- add sub-containers
- for prod - global error logging
- refactor som commands use player and some charobj
- refactor some commands use camelcase playerId and some underscore player_id
- refactor game object to have characteristics like pickable to influence the behaviour, rather then decribing everything manually
- fix bug that you can build where you stand and then cannot move