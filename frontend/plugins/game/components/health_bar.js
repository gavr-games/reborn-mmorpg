import * as BABYLON from 'babylonjs'

const HealthBarWidth = 2.0
const HealthBarHeight = 0.2
const HealthBarAlpha = 0.6

class HealthBar {
  constructor (value, maxValue, mesh, scene) {
    this.scene = scene
    this.value = value
    this.maxValue = maxValue
    const pos = mesh.position

    // draw healthbar plane
    this.plane = BABYLON.MeshBuilder.CreatePlane('healthbar-plane', { height: HealthBarHeight, width: HealthBarWidth }, scene)
    this.plane.position.x = pos.x
    this.plane.position.z = pos.z
    this.plane.position.y = mesh.getHierarchyBoundingVectors().max.y + HealthBarHeight * 2
    this.plane.material = new BABYLON.StandardMaterial('HealthBarMaterial', scene)
    this.plane.material.diffuseColor.set(1, 0, 0)
    this.plane.material.alpha = HealthBarAlpha
    this.plane.convertToUnIndexedMesh()
    this.plane.doNotSyncBoundingInfo = true
    this.plane.isPickable = false
    this.updateWidth()
  }

  update (value, maxValue, pos) {
    this.value = value
    this.maxValue = maxValue
    this.updateWidth()
    this.plane.position.x = pos.x
    this.plane.position.z = pos.z
    if (this.scene.cameras[0]) {
      this.plane.rotation = this.scene.cameras[0].rotation
    }
  }

  updateWidth () {
    this.plane.scaling = new BABYLON.Vector3(this.value / this.maxValue, 1, 1)
  }

  remove () {
    if (this.plane) {
      this.plane.dispose()
    }
    this.plane = null
  }
}

export default HealthBar
