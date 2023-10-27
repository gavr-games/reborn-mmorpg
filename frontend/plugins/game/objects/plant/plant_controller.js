import PlantObserver from "~/plugins/game/objects/plant/plant_observer";
import PlantState from "~/plugins/game/objects/plant/plant_state";
import { EventBus } from "~/plugins/game/event_bus";

class PlantController {
  constructor(gameObject) {
    this.state = new PlantState(gameObject);
    this.observer = new PlantObserver(this.state);
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default PlantController;
