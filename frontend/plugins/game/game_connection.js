import { EventBus } from "~/plugins/game/event_bus";
import GameObserver from "~/plugins/game/game_observer";

class GameConnnection {
  constructor() {
    this.conn = null
  }

  init(token, character_id) {
    if (window["WebSocket"]) {
      this.conn = new WebSocket("ws://" + document.location.host + "/engine/ws?token=" + token + "&character_id=" + character_id);
      this.conn.onclose = function (evt) {
        console.log("Engine ws connection is closed")
        window.location.href = "/login"
      };
      this.conn.onmessage = function (evt) {
        const messages = evt.data.split("\n")
        if (GameObserver.loaded) {
          GameObserver.scene.blockfreeActiveMeshesAndRenderingGroups = true
        }
        messages.forEach(message => {
          const data = JSON.parse(message);
          switch(data["ResponseType"]) {
            case "update_object":
              EventBus.$emit(data["ResponseType"], data["ResponseData"]["object"])
              break;
            case "remove_object":
              EventBus.$emit(data["ResponseType"], data["ResponseData"]["object"])
              break;
            case "add_object":
              EventBus.$emit(data["ResponseType"], data["ResponseData"]["object"])
              break;
            default:
              EventBus.$emit(data["ResponseType"], data["ResponseData"])
          }
        });
        if (GameObserver.loaded) {
          GameObserver.scene.blockfreeActiveMeshesAndRenderingGroups = false
        }
      };
    }
  }

  sendCmd(cmd, params) {
    if (this.conn) {
      let msg = JSON.stringify({
        cmd: cmd,
        params: params
      })
      this.conn.send(msg)
    }
  }
}

const gameConnnection = new GameConnnection();

export default gameConnnection;
