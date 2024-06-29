# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## Roadmap for Demo
- Add dragon breeding to get new eggs
  - breeding cave
  - teleport dragons for breeding to cave
  - DNA mutation system based on vitamins and food
    - each type of dragons requires BASIC set of vitamins just to make breeding possible
    - each pair of dragons can produce set of dragon eggs with probability once in a period of time
    - you can increase probability by feeding dragons the right way
    - you need a set of vitamins to increase probability breeding certain egg
    - you can feed dragons more, so they get more vitamins and increase probability X times per set
    - all probabilities are combined and you get weights
    - deduct vitamins for breeding.
- Dragon special abilities
- Improve path finding
- improve ALL visuals (https://store.steampowered.com/app/1726130/Pathless_Woods/)
  - refactor UI
  - delayed actions for pickup and others
  - default actions - select target, pickup
  - make craft in real world easier
  - add animations and effects
  - replace ugly models
  - effect of camera acceleration and braking
  - More buildings and furniture for houses
    - doors
    - normal walls

## TODO
- BUG: dungeon keeper not available after coming back from dungeon
- BUG: deselect target, when killed by dragons
- BUG: monster ghosts stay after death
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
