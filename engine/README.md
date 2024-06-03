# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## Roadmap for Demo
- More crops and food to raise dragons
  - fishing
  - fire
  - cook fish
- Add dragon breeding to get new eggs
  - breeding cave
  - teleport dragons for breeding to cave
  - DNA mutation system based on vitamins and food
- Game Master
  - add building
  - modify object JSON
- improve ALL visuals
  - refactor UI
  - add animations and effects
  - replace ugly models
  - More buildings and furniture for houses
    - doors
    - table
    - normal walls

## TODO
- BUG: dungeon keeper not available after back  from dungeon
- BUG: deselect target, when killed by dragons
- add Transactions (Mutexes) when objects cnahge
- change game atlas from type/kind to tags
- add leveling skills fot character and skills requirements for craft
- add more complex leveling formula, which takes into account many parameters like mob lvl and etc
- add more combat mechanics like defence, inasion, resists, etc.
- dragons get special abilities for feeding and vitamins to grow
- add path finding for mobs
- add database of game objects, craft OR editor to manage atlases (move atlases to json files where possible)
- switch from json to protobuf for server-client communication
- for prod - global error logging
- refactor some commands use player and some charobj
- refactor some commands use camelcase playerId and some underscore player_id
- refactor game object to have characteristics like pickable to influence the behaviour, rather then decribing everything manually
- think about sending only what is changed in game object
- FIX: login after disconect works only from second try
- Migrate test to https://onsi.github.io/ginkgo/#getting-started
- Cover engine with tests
- Add golang building cache to volume? speedup test and first launch