class RockState {
  constructor(gameObject) {
    this.x = gameObject["Properties"]["x"]
    this.y = gameObject["Properties"]["y"]
    this.id = gameObject["Id"]
    this.kind = gameObject["Properties"]["kind"]
    this.payload = gameObject;
  }
}

export default RockState;
