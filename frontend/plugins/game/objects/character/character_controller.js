import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";
import { EventBus } from "~/plugins/game/event_bus";

class CharacterController {
  constructor(gameObject, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.state = new CharacterState(gameObject)
    this.observer = new CharacterObserver(this.state, this.myCharacterId)
    // Display character target
    if (this.state.player_id == this.myCharacterId && this.state.payload.Properties["target"]) {
      EventBus.$emit("select_target", this.state.payload.Properties["target"])
    }
  }

  update(gameObject) {
    this.state.update(gameObject)
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default CharacterController;
