import { EventBus } from "~/plugins/game/event_bus";
import GameObserver from "~/plugins/game/game_observer";
import Border from "~/plugins/game/components/border";

class ClaimAreaObserver {
  constructor(state) {
    this.scene = null;
    this.state = state;
    this.border = null;
    this.meshRotation = 0
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene;
      this.create();
    } else {
      EventBus.$on("scene-created", scene => {
        this.scene = scene;
        this.create();
      });
    }

  }

  create() {
    this.border = new Border(this.state.x, this.state.y, this.state.width, this.state.height, this.scene)
  }

  remove() {
    EventBus.$off("scene-created", this.sceneCreatedCallback)
    if (this.border) {
      this.border.remove()
      this.border = null
    }
    this.state = null
  }
}

export default ClaimAreaObserver;
