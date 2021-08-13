import SurfaceObserver from "~/plugins/game/objects/surface/surface_observer";
import SurfaceState from "~/plugins/game/objects/surface/surface_state";
import { EventBus } from "~/plugins/game/event_bus";

class SurfaceController {
  constructor(gameObject) {
    this.state = new SurfaceState(gameObject);
    this.observer = new SurfaceObserver(this.state);
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

export default SurfaceController;
