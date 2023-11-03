import NpcObserver from "~/plugins/game/objects/npc/npc_observer";
import NpcState from "~/plugins/game/objects/npc/npc_state";
import { EventBus } from "~/plugins/game/event_bus";

class NpcController {
  constructor(gameObject) {
    this.state = new NpcState(gameObject);
    this.observer = new NpcObserver(this.state);
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default NpcController;
