# Engine

Start investigation of the code from `pkg/game/engine.go` the `Run()` func.
Later proceed to `pkg/game/engine/process_command.go` for possible actions in the game.

## Roadmap for Demo
- Body armor and Robe
- Add baby dragons and raising them
- Add dragon breeding to get new eggs
- Fighting with dragons in dungeons and leveling dragons
- More buildings and furniture for houses

## TODO
- add Transactions (Mutexes) when objects cnahge
- add leveling skills fot character and skills requirements for craft
- add more complex leveling formula, which takes into account many parameters like mob lvl and etc
- add database of game objects, craft OR editor to manage atlases (move atlases to json files where possible)
- switch from json to protobuf for server-client communication
- for prod - global error logging
- refactor some commands use player and some charobj
- refactor some commands use camelcase playerId and some underscore player_id
- refactor game object to have characteristics like pickable to influence the behaviour, rather then decribing everything manually
- think about sending only what is changed in game object
- FIX: login after disconect works only from second try
- Refactor: check where we can use CreateGameObject instead of CreateFromTemplate
- Migrate test to https://onsi.github.io/ginkgo/#getting-started
- Cover engine with tests