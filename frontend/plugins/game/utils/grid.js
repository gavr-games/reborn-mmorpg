import * as BABYLON from "babylonjs";

class Grid {
  constructor(scene) {
    this.scene = scene;
    this.width = 200;
    this.lines = []
  }

  create() {
    for (let n = 0; n < this.width; n++) {
      let myPoints = [
        new BABYLON.Vector3(0.5, 0.1, 1.0 * n - 0.5),
        new BABYLON.Vector3(
          this.width,
          0.1,
          1.0 * n - 0.5
        )
      ];
      this.lines.push(BABYLON.MeshBuilder.CreateLines(
        "lines",
        { points: myPoints },
        this.scene
      ))
    }
    for (let n = 0; n < this.width; n++) {
      let myPoints = [
        new BABYLON.Vector3(1.0 * n + 0.5, 0.1, -0.5),
        new BABYLON.Vector3(
          1.0 * n + 0.5,
          0.1,
          this.width
        )
      ];
      this.lines.push(BABYLON.MeshBuilder.CreateLines(
        "lines",
        { points: myPoints },
        this.scene
      ))
    }
  }

  remove() {
    this.lines.forEach((line) => {
      line.dispose()
    })
    this.lines = []
  }
}

export default Grid;
