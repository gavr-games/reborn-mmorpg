import * as BABYLON from 'babylonjs'
import { EventBus } from '~/plugins/game/event_bus'
import Atlas from '~/plugins/game/atlas/atlas'
import GameObserver from '~/plugins/game/game_observer'

class NpcObserver {
  constructor (state) {
    this.scene = null
    this.state = state
    this.container = null
    this.mesh = null
    this.meshRotation = Math.PI / 2
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
    this.container = Atlas.get(this.state.kind + 'Npc').instantiateModelsToScene()
    this.playAnimation('Idle')
    const mesh = this.container.rootNodes[0]
    mesh.setParent(null)
    mesh.name = 'npc-' + this.state.id
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    if (this.state.rotation) {
      const rotationDelta = this.meshRotation - this.state.rotation
      if (rotationDelta !== 0) {
        this.meshRotation = this.state.rotation
        mesh.rotate(BABYLON.Axis.Y, rotationDelta)
      }
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
  }

  remove () {
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    this.mesh.dispose()
    this.mesh = null
    this.state = null
  }

  playAnimation (name, loop = true) {
    if (this.container && this.currentAnimation !== name) {
      this.container.animationGroups.forEach((ag) => {
        if (ag.name.includes(name)) {
          ag.start(loop)
          this.currentAnimation = name
          if (!loop) {
            ag.onAnimationEndObservable.addOnce(() => {
              this.currentAnimation = 'Idle'
            })
          }
        } else {
          ag.reset()
          ag.stop()
        }
      })
    }
  }
}

export default NpcObserver
