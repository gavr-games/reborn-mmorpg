import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";
import { EventBus } from "~/plugins/game/event_bus";

class CharacterController {
  constructor(gameObject, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.state = new CharacterState(gameObject);
    this.observer = new CharacterObserver(this.state, this.myCharacterId);
    this.updateHandler = (gameObject) => {
      this.state.update(gameObject)
    };
    this.removeHandler = () => {
      this.remove()
    };
    EventBus.$on(`update_object_${this.state.id}`, this.updateHandler);
    EventBus.$on(`remove_object_${this.state.id}`, this.removeHandler);
  }

  remove() {
    EventBus.$off(`update_object_${this.state.id}`, this.updateHandler);
    EventBus.$off(`remove_object_${this.state.id}`, this.removeHandler);
    this.state = null
    this.observer.remove()
  }
}

export default CharacterController;
