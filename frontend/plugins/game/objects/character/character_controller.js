import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";
import { EventBus } from "~/plugins/game/event_bus";

class CharacterController {
  constructor(gameObject, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.state = new CharacterState(gameObject)
    this.observer = new CharacterObserver(this.state, this.myCharacterId)
    this.performMeleeHit = data => {
      if (data.object.Id == this.state.id) {
        this.observer.meleeHit(data.weapon)
      }
    };
    EventBus.$on("melee_hit_attempt", this.performMeleeHit)
    // Display character target
    if (this.state.player_id == this.myCharacterId && this.state.payload.Properties["target"]) {
      EventBus.$emit("select_target", this.state.payload.Properties["target"])
    }
    // Publish info about player's character
    if (this.state.player_id == this.myCharacterId) {
      EventBus.$emit("my-character-info", this.state.payload)
    }
  }

  update(gameObject) {
    this.state.update(gameObject)
  }

  remove() {
    EventBus.$off("melee_hit_attempt", this.performMeleeHit)
    this.state = null
    this.observer.remove()
  }

  selectAsTarget() {
    this.observer.selectAsTarget()
  }

  deselectAsTarget() {
    this.observer.deselectAsTarget()
  }
}

export default CharacterController;
