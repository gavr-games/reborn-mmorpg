import * as BABYLON from 'babylonjs'
import { EventBus } from '~/plugins/game/event_bus'
import Atlas from '~/plugins/game/atlas/atlas'
import GameObserver from '~/plugins/game/game_observer'

class SurfaceObserver {
  constructor (state) {
    this.scene = null
    this.state = state
    this.container = null
    this.mesh = null
    this.groun = null
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene
      this.create()
    } else {
      EventBus.$on('scene-created', (scene) => {
        this.scene = scene
        this.create()
      })
    }
  }

  create () {
    const mesh = Atlas.get(this.state.kind + 'Surface').createInstance('surface-' + this.state.id)
    mesh.setParent(null)
    mesh.isPickable = false
    mesh.name = 'surface-' + this.state.id
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    if (this.state.rotation) {
      mesh.rotate(BABYLON.Axis.Y, Math.PI / 2)
    }
    mesh.metadata = {
      x: this.state.x,
      y: this.state.y,
      id: this.state.id,
      state: this.state
    }
    mesh.setEnabled(true)
    mesh.freezeWorldMatrix()
    mesh.doNotSyncBoundingInfo = true
    this.mesh = mesh
    if (this.state.kind === 'water') {
      // Create ground (bottom for water)
      const groundMaterial = new BABYLON.StandardMaterial('groundMaterial', this.scene)
      groundMaterial.disableLighting = true
      groundMaterial.backFaceCulling = false
      groundMaterial.emissiveColor = new BABYLON.Color3(0, 0, 1)

      this.ground = BABYLON.MeshBuilder.CreateGround(`ground-${this.state.id}`, { height: 1, width: 1 }, this.scene)
      this.ground.position.x = this.state.x
      this.ground.position.y = -0.1
      this.ground.position.z = this.state.y
      this.ground.doNotSyncBoundingInfo = true
      this.ground.isPickable = false
      this.ground.freezeWorldMatrix()
      this.ground.material = groundMaterial
      this.ground.material.freeze()
      this.mesh.material.addToRenderList(this.ground)
    }
  }

  remove () {
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    if (this.mesh) {
      this.mesh.dispose()
    }
    if (this.ground) {
      this.ground.dispose()
    }
    this.mesh = null
    this.ground = null
    this.state = null
  }
}

export default SurfaceObserver
