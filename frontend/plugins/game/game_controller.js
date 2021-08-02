import GameObserver from "~/plugins/game/game_observer";
import ChatController from "~/plugins/game/chat/chat_controller";
import { EventBus } from "~/plugins/game/event_bus";
import SurfaceController from "./objects/surface/surface_controller";
import CharacterController from "./objects/character/character_controller";

class GameController {
  constructor() {
    this.conn = null
    this.gameObjects = []
    this.token = null
    this.characterId = null
    this.controls = {
      w: false,
      a: false,
      s: false,
      d: false,
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
    EventBus.$on("init_game", this.initGameObjectsHandler);
    EventBus.$on("keyup", this.keyUpHandler);
    EventBus.$on("keydown", this.keyDownHandler);
  }

  init(token, character_id) {
    this.token = token
    this.characterId = character_id
    if (window["WebSocket"]) {
      this.conn = new WebSocket("ws://" + document.location.host + "/engine/ws?token=" + token + "&character_id=" + character_id);
      this.conn.onclose = function (evt) {
        console.log("Engine ws connection is closed")
        window.location.href = "/login"
      };
      this.conn.onmessage = function (evt) {
        const data = JSON.parse(evt.data);
        EventBus.$emit(data["ResponseType"], data["ResponseData"])
      };
    }
    GameObserver.init();
    ChatController.init(token, character_id);
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

  handleDownUp() {
    if (['w', 'a', 's', 'd'].includes(key)) {
      this.controls[key] = true
      this.moveCharacter()
    }
  }

  moveCharacter() {
    if (!this.controls.w && !this.controls.a &&
      !this.controls.s && !this.controls.d) {
        //Send stop
        return
      }
    //if
  }

  destroy() {
    ChatController.destroy();
  }
}

const gameController = new GameController();

export default gameController;
