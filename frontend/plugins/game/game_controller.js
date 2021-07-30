import GameObserver from "~/plugins/game/game_observer";
import ChatController from "~/plugins/game/chat/chat_controller";
import { EventBus } from "~/plugins/game/event_bus";

class GameController {
  constructor() {
    this.conn = null
  }

  init(token, character_id) {
    GameObserver.init();
    ChatController.init(token, character_id);
    if (window["WebSocket"]) {
      this.conn = new WebSocket("ws://" + document.location.host + "/engine/ws?token=" + token + "&character_id=" + character_id);
      this.conn.onclose = function (evt) {
        console.log("Engine ws connection is closed")
      };
      this.conn.onmessage = function (evt) {
        console.log(evt.data);
        EventBus.$emit(evt.data.ResponseType, evt.data.ResponseData)
      };
    }
  }

  destroy() {
    ChatController.destroy();
  }
}

const gameController = new GameController();

export default gameController;
