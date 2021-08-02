import { EventBus } from "~/plugins/game/event_bus";

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
        const data = JSON.parse(evt.data);
        EventBus.$emit(data["ResponseType"], data["ResponseData"])
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
