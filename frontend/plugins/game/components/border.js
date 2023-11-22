import * as BABYLON from "babylonjs"

const Y = 0.1

class Border {
  constructor(x, y, width, height, scene) {
    this.scene = scene
    this.lines = []
    let points = [
      new BABYLON.Vector3(x - width / 2, Y, y + height / 2),
      new BABYLON.Vector3(x + width / 2, Y, y + height / 2),
      new BABYLON.Vector3(x + width / 2, Y, y - height / 2),
      new BABYLON.Vector3(x - width / 2, Y, y - height / 2),
      new BABYLON.Vector3(x - width / 2, Y, y + height / 2),
    ];
    this.lines.push(BABYLON.MeshBuilder.CreateLines(
      "claimLines",
      { points: points, closed: true },
      this.scene
    ))
    this.lines.color = new BABYLON.Color3(1, 0, 0);
  }

  remove() {
    this.lines.forEach((line) => {
      line.dispose()
    })
    this.lines = []
  }
}

export default Border;
