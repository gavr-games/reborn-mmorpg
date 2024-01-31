import * as BABYLON from 'babylonjs'

const ALPHA = 1.0
const Y = 0.1

class MeleeHitArea {
  constructor (radius, angle, cooldown, rotation, pos, scene) {
    this.scene = scene
    this.disc = BABYLON.MeshBuilder.CreateDisc('hitarea-disc', { radius, arc: angle / 360, tessellation: 12 }, scene)
    this.disc.position.x = pos.x
    this.disc.position.z = pos.z
    this.disc.position.y = Y
    this.disc.rotate(BABYLON.Axis.X, Math.PI / 2)
    this.disc.rotate(BABYLON.Axis.Z, rotation - Math.PI / 4)
    this.disc.material = new BABYLON.StandardMaterial('MeleeHitAreaMaterial', scene)
    this.disc.material.diffuseColor.set(1, 0, 0)
    this.disc.material.alpha = ALPHA
    this.disc.convertToUnIndexedMesh()
    this.disc.doNotSyncBoundingInfo = true
    this.disc.isPickable = false

    setTimeout(() => {
      this.remove()
    }, cooldown)
    this.interval = setInterval(() => {
      this.disc.material.alpha = this.disc.material.alpha - 0.1
    }, cooldown / 10)
  }

  remove () {
    if (this.disc) {
      this.disc.dispose()
      clearInterval(this.interval)
    }
    this.disc = null
  }
}

export default MeleeHitArea
