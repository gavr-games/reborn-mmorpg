# Engine Mechanics

## Delayed Actions
Some game actions, like craft/chop/chip take time. So before calling a function `Craft`, `Chop`, `Chip` the game need actually to wait for a specific ammount of time.

To do that the engine loops through all game objects and looks if a particular GO has `CurrentAction`.

CurrentAction actually means a function to execute in the future.

To store function to be executed engine uses [DelayedAction struct](../engine/pkg/game/entity/delayed_action.go).
This struct has FuncName to be executed and time to wait (`TimeLeft`).

Engine tick by tick decreases `TimeLeft`. When it becomes 0, the Engine goes to [Atlas](../engine/pkg/game/engine/craft/atlas.go), finds the appropriate golang function and executes it with params.

## Move To Coords
This is an engine mechanism to make character or mob to automatically move to some coords or approach and object.
For example walk to the tree to chop it.

This has higher priority then DelayedActions mechanism above. So character First moves to coords, then does Delayed Actions. This allows to do things like approach a tree and then chop it.

Each GO stores where to move in struct's field called `moveToCoords`.

You can ask character two move to coords in 2 modes: 1 - move to exact coords (character will try to reach the exact coords), this is used for mobs to follow the owner; 2 - move close to bounds, this is used to move character close to object but don't intersect with it (move close to object and build a wall).
