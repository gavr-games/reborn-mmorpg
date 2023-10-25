import * as BABYLON from "babylonjs"

const ALPHA = 0.8
const Y = 0.1

class MeleeHitArea {
  constructor(radius, angle, cooldown, rotation, pos, scene) {
    this.scene = scene
    this.disc = BABYLON.MeshBuilder.CreateDisc("hitarea-disc", {radius: radius, arc: angle / 360, tessellation: 12}, scene)
    this.disc.position.x = pos.x
    this.disc.position.z = pos.z
    this.disc.position.y = Y
    this.disc.rotate(BABYLON.Axis.X, Math.PI / 2);
    this.disc.rotate(BABYLON.Axis.Z, rotation - Math.PI / 4);
    this.disc.material = new BABYLON.StandardMaterial("MeleeHitAreaMaterial", scene)
    this.disc.material.diffuseColor.set(1, 0, 0)
    this.disc.material.alpha = ALPHA;

    setTimeout(() => {
      this.remove()
    }, cooldown)
  }

  remove() {
    if (this.disc) {
      this.disc.dispose();
    }
    this.disc = null;
  }
}

export default MeleeHitArea;
