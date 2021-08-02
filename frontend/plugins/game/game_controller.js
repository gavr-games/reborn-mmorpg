import GameObserver from "~/plugins/game/game_observer";
import ChatController from "~/plugins/game/chat/chat_controller";
import { EventBus } from "~/plugins/game/event_bus";
import SurfaceController from "./objects/surface/surface_controller";
import CharacterController from "./objects/character/character_controller";
import GameConnnection from "./game_connection";

class GameController {
  constructor() {
    this.gameObjects = []
    this.token = null
    this.characterId = null
    this.controls = {
      w: false,
      a: false,
      s: false,
      d: false,
      lastCmd: ""
    }
    this.initGameObjectsHandler = gameObjects => {
      this.initGameObjects(gameObjects)
    };
    this.keyUpHandler = key => {
      this.handleKeyUp(key)
    };
    this.keyDownHandler = key => {
      this.handleKeyDown(key)
    };
    EventBus.$on("init_game", this.initGameObjectsHandler)
    EventBus.$on("keyup", this.keyUpHandler)
    EventBus.$on("keydown", this.keyDownHandler)
  }

  init(token, character_id) {
    this.token = token
    this.characterId = character_id

    GameConnnection.init(token, character_id)
    GameObserver.init()
    ChatController.init(token, character_id)
  }

  initGameObjects(gameObjects) {
    gameObjects["visible_objects"].forEach(gameObj => {
      switch(gameObj["Type"]) {
        case "surface":
          this.gameObjects[gameObj["Id"]] = new SurfaceController(gameObj)
        break;
        case "player":
          if (gameObj["Properties"]["player_id"] == this.characterId) {
            this.gameObjects[gameObj["Id"]] = new CharacterController(gameObj)
          }
        break;
      }
    });
  }

  handleKeyUp(key) {
    if (['w', 'a', 's', 'd'].includes(key)) {
      this.controls[key] = false
      this.moveCharacter()
    }
  }

  handleKeyDown(key) {
    if (['w', 'a', 's', 'd'].includes(key)) {
      this.controls[key] = true
      this.moveCharacter()
    }
  }

  moveCharacter() {
    let directionSum = 0
    if (this.controls.w) directionSum += 1
    if (this.controls.a) directionSum += 10
    if (this.controls.s) directionSum += 100
    if (this.controls.d) directionSum += 1000
    let cmd = "stop"
    switch (directionSum) {
      case 1:
        cmd = "move_north_west"
        break;
      case 10:
        cmd = "move_south_west"
        break;
      case 100:
        cmd = "move_south_east"
        break;
      case 1000:
        cmd = "move_north_east"
        break;
      case 11:
        cmd = "move_west"
        break;
      case 110:
        cmd = "move_south"
        break;
      case 1100:
        cmd = "move_east"
        break;
      case 1001:
        cmd = "move_north"
        break;
    }
    if (cmd != this.controls.lastCmd) {
      GameConnnection.sendCmd(cmd)
      this.controls.lastCmd = cmd
    }
  }

  destroy() {
    ChatController.destroy();
  }
}

const gameController = new GameController();

export default gameController;
