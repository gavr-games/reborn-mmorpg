import { EventBus } from "~/plugins/game/event_bus";

class CharacterMenuController {
  constructor() {
    this.gameConnection = null
    this.getCharacterInfo = message => {
      if (this.gameConnection) {
        this.gameConnection.sendCmd("get_character_info")
      }
    };
    EventBus.$on("get-character-info", this.getCharacterInfo);
  }

  init(gameConnection) {
    this.gameConnection = gameConnection
  }

  destroy() {
    EventBus.$off("get-character-info")
  }
}

const characterMenuController = new CharacterMenuController();

export default characterMenuController;
