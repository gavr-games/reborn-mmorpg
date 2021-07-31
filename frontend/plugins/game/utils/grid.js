import * as BABYLON from "babylonjs";

class Grid {
  constructor(scene) {
    this.scene = scene;
    this.width = 20;
  }

  create() {
    for (let n = 0; n < this.width; n++) {
      let myPoints = [
        new BABYLON.Vector3(0, 0.3, 1.0 * n),
        new BABYLON.Vector3(
          1.0 * this.width,
          0.3,
          1.0 * n
        )
      ];
      BABYLON.MeshBuilder.CreateLines(
        "lines",
        { points: myPoints },
        this.scene
      );
    }
    for (let n = 0; n < this.width; n++) {
      let myPoints = [
        new BABYLON.Vector3(1.0 * n, 0.3, 0),
        new BABYLON.Vector3(
          1.0 * n,
          0.3,
          1.0 * this.width
        )
      ];
      BABYLON.MeshBuilder.CreateLines(
        "lines",
        { points: myPoints },
        this.scene
      );
    }
  }
}

export default Grid;
