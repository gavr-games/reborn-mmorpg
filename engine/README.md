# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## TODO
- add claims
  - craft obelisk
  - make sure only 1 obelisk per player
  - teleport to obelisk
  - show obelisk area
  - don't allow other player to interract wih items on someones obelisk area
  - no collision with other obelisks or too close positions
  - obelisk rent
  - destroy obelisk if the rent is gone
- add baby dragons and raising them
- add dragon breeding to get new eggs
- add fighting and leveling dragins
- add leveling skills fot character and skills requirements for craft
- send visible objects updates once per 500ms and check performance
- add database of game objects, craft OR editor to manage atlases (move atlases to json files where possible)
- switch from json to protobuf for server-client communication
- add minimap feature
- add sub-containers
- for prod - global error logging
- refactor some commands use player and some charobj
- refactor some commands use camelcase playerId and some underscore player_id
- refactor game object to have characteristics like pickable to influence the behaviour, rather then decribing everything manually
- think about sending only what is changed in game object
- FIX: login after disconect works only from second try
- add GameMasters
- Refactor: check where we can use CreateGameObject instead of CreateFromTemplate
