# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## TODO
- send visible objects updates once per 500ms and check performance
- allow mobs to hit back
  - add mob following the target
  - add mob trying to hit
  - add character reborn (teleport and full hp)
- add healing potions
- don't allow to move out of the map
- add database of game objects, craft OR editor to manage atlases
- switch from json to protobuf
- add minimap feature
- add sub-containers
- for prod - global error logging
- refactor some commands use player and some charobj
- refactor some commands use camelcase playerId and some underscore player_id
- refactor game object to have characteristics like pickable to influence the behaviour, rather then decribing everything manually
- think about sending only what is changed in game object
- FIX: login after disconect works only from second try