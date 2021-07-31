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
    this.initGameObjectsHandler = gameObjects => {
      this.initGameObjects(gameObjects)
    };
    EventBus.$on("init_game", this.initGameObjectsHandler);
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

  destroy() {
    ChatController.destroy();
  }
}

const gameController = new GameController();

export default gameController;
