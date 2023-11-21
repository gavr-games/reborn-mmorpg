import ClaimAreaObserver from "~/plugins/game/objects/claim_area/claim_area_observer";
import ClaimAreaState from "~/plugins/game/objects/claim_area/claim_area_state";

class ClaimAreaController {
  constructor(gameObject) {
    this.state = new ClaimAreaState(gameObject);
    this.observer = new ClaimAreaObserver(this.state);
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default ClaimAreaController;
