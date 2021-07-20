import * as BABYLON from "babylonjs";

class Camera {
  constructor(scene, canvas) {
    this.scene = scene;
    this.canvas = canvas;
    this.camera = null;
  }

  create() {
    this.camera = new BABYLON.ArcRotateCamera(
      "MainCamera",
      Math.PI / 2,
      Math.PI / 4,
      5,
      new BABYLON.Vector3(
        1,
        1,
        1
      ),
      this.scene
    );
    this.camera.attachControl(this.canvas, true);
  }
}

export default Camera;
