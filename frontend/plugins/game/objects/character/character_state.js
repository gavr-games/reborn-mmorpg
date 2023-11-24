class CharacterState {
  constructor(gameObject) {
    this.update(gameObject)
  }

  update(gameObject) {
    // x and y transformation is required because in engine upper left corner is stored as x/y,
    // but on frontend all assets have pivot points in the center of the object.
    const transformX = gameObject["Rotation"] % (Math.PI / 2) == 0 ? gameObject["Properties"]["width"] / 2 : gameObject["Properties"]["height"] / 2
    const transformY = gameObject["Rotation"] % (Math.PI / 2) == 0 ? gameObject["Properties"]["height"] / 2 : gameObject["Properties"]["width"] / 2
    this.x = gameObject["Properties"]["x"] + transformX
    this.y = gameObject["Properties"]["y"] + transformY
    this.rotation = gameObject["Rotation"]
    this.id = gameObject["Id"]
    this.speed = gameObject["Properties"]["speed"]
    this.speed_x = gameObject["Properties"]["speed_x"]
    this.speed_y = gameObject["Properties"]["speed_y"]
    this.player_id = gameObject["Properties"]["player_id"]
    this.health = gameObject["Properties"]["health"]
    this.max_health = gameObject["Properties"]["max_health"]
    this.payload = gameObject;
  }
}

export default CharacterState;
