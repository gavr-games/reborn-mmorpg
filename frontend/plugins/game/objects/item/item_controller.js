import ItemObserver from "~/plugins/game/objects/item/item_observer";
import ItemState from "~/plugins/game/objects/item/item_state";
import { EventBus } from "~/plugins/game/event_bus";

class ItemController {
  constructor(gameObject) {
    this.state = new ItemState(gameObject);
    this.observer = new ItemObserver(this.state);
  }

  update(gameObject) {
    this.state.update(gameObject)
    this.observer.updateLifted()
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default ItemController;
