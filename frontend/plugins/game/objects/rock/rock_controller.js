import RockObserver from "~/plugins/game/objects/rock/rock_observer";
import RockState from "~/plugins/game/objects/rock/rock_state";
import { EventBus } from "~/plugins/game/event_bus";

class RockController {
  constructor(gameObject) {
    this.state = new RockState(gameObject);
    this.observer = new RockObserver(this.state);
    this.removeHandler = () => {
      this.remove()
    };
    EventBus.$on(`remove_object_${this.state.id}`, this.removeHandler);
  }

  remove() {
    EventBus.$off(`remove_object_${this.state.id}`, this.removeHandler);
    this.state = null
    this.observer.remove()
  }
}

export default RockController;
