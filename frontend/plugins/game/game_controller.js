import GameObserver from "~/plugins/game/game_observer";

class GameController {
  init() {
    GameObserver.init();
  }
}

const gameController = new GameController();

export default gameController;
