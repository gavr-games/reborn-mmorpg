import GameObserver from "~/plugins/game/game_observer";
import ChatController from "~/plugins/game/chat/chat_controller";

class GameController {
  init(token, character_id) {
    GameObserver.init();
    ChatController.init(token, character_id);
  }

  destroy() {
    ChatController.destroy();
  }
}

const gameController = new GameController();

export default gameController;
