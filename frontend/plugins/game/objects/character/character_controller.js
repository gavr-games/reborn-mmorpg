import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";

class CharacterController {
  constructor(gameObject) {
    this.state = new CharacterState(gameObject);
    this.observer = new CharacterObserver(this.state);
  }
}

export default CharacterController;
