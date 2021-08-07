class CharacterState {
  constructor(gameObject) {
    this.update(gameObject)
  }

  update(gameObject) {
    this.x = gameObject["Properties"]["x"]
    this.y = gameObject["Properties"]["y"]
    this.id = gameObject["Id"]
    this.speed = gameObject["Properties"]["speed"]
    this.speed_x = gameObject["Properties"]["speed_x"]
    this.speed_y = gameObject["Properties"]["speed_y"]
    this.player_id = gameObject["Properties"]["player_id"]
    this.payload = gameObject;
  }
}

export default CharacterState;
