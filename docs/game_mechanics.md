# Game Mechanics

## Claims

Players can build bases/farms. Honestly speaking they can build it anywhere. 

You need a base to protect your items. The items are protected from other players (cannot take/destroy/chop) only if you build something inside the Claim Area.

Claim Area - is an area aroun Claim Obelisk. you can craft it from craft menu.

Claim Obelisk can expire after 4 weeks. To extend the rent you should pay for it.

Each character can have only one obelisk installed.

Two Claim Areas cannot intersect.

## Containers

Containers have different size. You can put smaller containers inside bigger ones.

## Collision

Each Game Object is either `collidable` or not. If so it means moving characters cannot go through it. For example: character cannot go through the wall, but can go through a mob.
In addition to that there is a `craft_collidable` property. If set to true, players will be able to craft something in real world, even if the crafted object collides with `"craft_collidable": false` object. For example: players can craft building on claim area.
