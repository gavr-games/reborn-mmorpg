import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";

class CharacterController {
  constructor(gameObject, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.state = new CharacterState(gameObject);
    this.observer = new CharacterObserver(this.state, this.myCharacterId);
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
