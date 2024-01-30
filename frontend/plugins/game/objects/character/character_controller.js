import CharacterObserver from "./character_observer";
import CharacterState from "./character_state";
import { EventBus } from "~/plugins/game/event_bus";

class CharacterController {
  constructor(gameObject, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.state = new CharacterState(gameObject)
    this.observer = new CharacterObserver(this.state, this.myCharacterId)
    this.emulateTimeout = null
    this.performMeleeHit = data => {
      if (data.object.Id == this.state.id) {
        this.observer.meleeHit(data.weapon)
      }
    }
    // This helps to predict character movement before server answer and get rid of "jumps" (ping)
    this.emulateMove = data => {
      if (this.emulateTimeout !== null) {
        clearTimeout(this.emulateTimeout)
      }
      this.emulateTimeout = setTimeout(() => {
        if (Date.now() - data.ping > this.state.lastUpdate) { // there was no answer from server yet
          this.setXYSpeeds(data.cmd)
        }
        this.emulateTimeout = null
      }, data.ping)
    }
    EventBus.$on("melee_hit_attempt", this.performMeleeHit)
    // Display character target
    if (this.state.player_id == this.myCharacterId && this.state.payload.Properties["target"]) {
      EventBus.$emit("select_target", this.state.payload.Properties["target"])
    }
    if (this.state.player_id == this.myCharacterId) {
      EventBus.$emit("my-character-info", this.state.payload) // Publish info about player's character
      EventBus.$on("emulate-character-move", this.emulateMove)
    }
  }

  update(gameObject) {
    this.state.update(gameObject)
  }

  remove() {
    EventBus.$off("melee_hit_attempt", this.performMeleeHit)
    EventBus.$off("emulate-character-move", this.emulateMove)
    this.state = null
    this.observer.remove()
  }

  selectAsTarget() {
    this.observer.selectAsTarget()
  }

  deselectAsTarget() {
    this.observer.deselectAsTarget()
  }

  setXYSpeeds(direction) {
    const speed = this.state.speed
	  const axisSpeed = Math.sqrt(speed * speed / 2)

    switch (direction) {
      case "move_north":
        this.state.speed_x = 0
        this.state.speed_y = speed
        break;
      case "move_south":
        this.state.speed_x = 0
        this.state.speed_y = -speed
        break;
      case "move_east":
        this.state.speed_x = speed
        this.state.speed_y = 0
        break;
      case "move_west":
        this.state.speed_x = -speed
        this.state.speed_y = 0
        break;
      case "move_north_east":
        this.state.speed_x = axisSpeed
        this.state.speed_y = axisSpeed
        break;
      case "move_north_west":
        this.state.speed_x = -axisSpeed
        this.state.speed_y = axisSpeed
        break;
      case "move_south_east":
        this.state.speed_x = axisSpeed
        this.state.speed_y = -axisSpeed
        break;
      case "move_south_west":
        this.state.speed_x = -axisSpeed
        this.state.speed_y = -axisSpeed
        break;
      case "stop":
        this.state.speed_x = 0
        this.state.speed_y = 0
        break;
    }
  }
}

export default CharacterController;
