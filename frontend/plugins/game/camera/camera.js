import * as BABYLON from "babylonjs";

class Camera {
  constructor(scene, canvas, character) {
    this.scene = scene;
    this.canvas = canvas;
    this.character = character;
    this.camera = null;
  }

  create() {
    this.camera = new BABYLON.FollowCamera(
      "MainCamera",
      new BABYLON.Vector3(
        this.character.mesh.position.x + 20,
        this.character.mesh.position.y + 20,
        this.character.mesh.position.z - 20
      ),
      this.scene
    );

    var ground = BABYLON.Mesh.CreatePlane("g", 1, this.scene);
    ground.position = new BABYLON.Vector3(0.5, 0, 0.5);
    ground.rotation.x = Math.PI / 2;

    this.camera.cameraAcceleration = 0; // how fast to move
    this.camera.maxCameraSpeed = 20; // speed limit
    this.camera.attachControl(this.canvas, true);

    this.camera.lockedTarget = this.character.mesh;
  }
}

export default Camera;
