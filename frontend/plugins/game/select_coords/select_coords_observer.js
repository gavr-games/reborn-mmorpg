import * as BABYLON from 'babylonjs'
import { EventBus } from '~/plugins/game/event_bus'
import Atlas from '~/plugins/game/atlas/atlas'
import GameObserver from '~/plugins/game/game_observer'
import addAlpha from '~/plugins/game/utils/add_alpha'

class SelectCoordsObserver {
  constructor () {
    this.scene = null
    this.container = null
    this.mesh = null
    this.rotation = 0
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene
    } else {
      EventBus.$on('scene-created', (scene) => {
        this.scene = scene
      })
    }
  }

  create (itemKey, x, y) {
    let mesh = null
    try {
      mesh = Atlas.get(itemKey + 'Item').clone('select-coords-item')
    } catch (e) {
      mesh = Atlas.get(itemKey + 'Plant').clone('select-coords-item')
    }
    mesh.setParent(null)
    mesh.setEnabled(true)
    mesh.name = 'select-coords-item'
    mesh.position.x = x
    mesh.position.y = 0
    mesh.position.z = y
    mesh.setEnabled(true)
    mesh.doNotSyncBoundingInfo = true
    mesh.isPickable = false
    this.mesh = mesh
    addAlpha(this.mesh, this.scene, 0.7)
    GameObserver.grid.create()
  }

  update (x, y) {
    this.mesh.position.x = x
    this.mesh.position.z = y
  }

  rotate () {
    let angle = Math.PI / 2
    if (this.rotation === 1) {
      angle = -Math.PI / 2
    }
    this.rotation = this.rotation === 0 ? 1 : 0
    this.mesh.rotate(BABYLON.Axis.Y, angle)
  }

  remove () {
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    if (this.mesh) {
      this.mesh.dispose()
      this.mesh = null
    }
    GameObserver.grid.remove()
  }
}

export default SelectCoordsObserver
