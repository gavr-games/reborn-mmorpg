import * as BABYLON from "babylonjs"

const HealthBarWidth = 2.0
const HealthBarHeight = 0.5
const HealthBarAlpha = 0.6

class HealthBar {
  constructor(value, maxValue, pos, camera, scene) {
    this.scene = scene
    this.value = value
    this.maxValue = maxValue

    //draw healthbar plane
    this.plane = BABYLON.MeshBuilder.CreatePlane("healthbar-plane", {height: HealthBarHeight, width: HealthBarWidth}, scene)
    this.plane.position.x = pos.x
    this.plane.position.z = pos.z
    this.plane.rotation = camera.rotation
    this.plane.material = new BABYLON.StandardMaterial("HealthBarMaterial", scene)
    this.plane.material.diffuseColor.set(1, 0, 0)
    this.plane.material.alpha = HealthBarAlpha;
    this.updateWidth()
  }

  update(value, maxValue, pos) {
    this.value = value
    this.maxValue = maxValue
    this.updateWidth()
    this.updateWidth
    this.plane.position.x = pos.x
    this.plane.position.z = pos.z
  }

  updateWidth() {
    this.plane.width = HealthBarWidth * this.value / this.maxValue
  }

  remove() {
    if (this.plane) {
      this.plane.dispose();
    }
    this.plane = null;
  }
}

export default HealthBar;
