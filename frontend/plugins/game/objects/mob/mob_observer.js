import * as BABYLON from 'babylonjs'
import { EventBus } from '~/plugins/game/event_bus'
import Atlas from '~/plugins/game/atlas/atlas'
import GameObserver from '~/plugins/game/game_observer'
import HealthBar from '~/plugins/game/components/health_bar'
import MeleeHitArea from '~/plugins/game/components/melee_hit_area'
import HighlightShape from '~/plugins/game/components/highlight_shape'
import freezeMaterials from '~/plugins/game/utils/freeze_materials'

class MobObserver {
  constructor (state) {
    this.scene = null
    this.state = state
    this.container = null
    this.mesh = null
    this.meshRotation = Math.PI / 2
    this.currentAnimation = null
    this.healthbar = null
    this.targetHighlight = null
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
    this.container = Atlas.get(this.state.kind + 'Mob').instantiateModelsToScene()
    this.playAnimation('Idle')
    const mesh = this.container.rootNodes[0]
    mesh.setParent(null)
    freezeMaterials(mesh, this.scene)
    mesh.name = 'mob-' + this.state.id
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
    mesh.doNotSyncBoundingInfo = true
    this.mesh = mesh
    this.healthbar = new HealthBar(this.state.health, this.state.max_health, this.mesh.position, this.scene)
    GameObserver.addRenderObserver(`mob-${this.state.id}`, this)
  }

  update (renderInterval) {
    if (this.state.speed_x !== 0 || this.state.speed_y !== 0) {
      this.playAnimation('Walk')
      this.state.x = this.state.x + this.state.speed_x / 1000 * renderInterval
      this.state.y = this.state.y + this.state.speed_y / 1000 * renderInterval
      const rotationAngle = Math.atan2(
        this.state.speed_y,
        this.state.speed_x
      )
      const rotationDelta = this.meshRotation - rotationAngle
      if (rotationDelta !== 0) {
        this.meshRotation = rotationAngle
        this.mesh.rotate(BABYLON.Axis.Y, rotationDelta)
      }
    } else {
      this.playAnimation('Idle')
    }
    this.mesh.position.x = this.state.x
    this.mesh.position.z = this.state.y
    this.healthbar.update(this.state.health, this.state.max_health, this.mesh.position)
    if (this.targetHighlight) {
      this.targetHighlight.update(this.mesh.position)
    }
  }

  remove () {
    GameObserver.removeRenderObserver(`mob-${this.state.id}`)
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    if (this.healthbar) {
      this.healthbar.remove()
      this.healthbar = null
    }
    if (this.targetHighlight) {
      this.targetHighlight.remove()
      this.targetHighlight = null
    }
    if (this.mesh !== null) {
      this.mesh.dispose()
    }
    this.mesh = null
    this.state = null
  }

  selectAsTarget () {
    if (this.mesh) {
      this.targetHighlight = new HighlightShape(this.state.payload, this.mesh.position, this.scene)
    } else {
      setTimeout(() => {
        this.selectAsTarget()
      }, 100)
    }
  }

  deselectAsTarget () {
    if (this.targetHighlight) {
      this.targetHighlight.remove()
      this.targetHighlight = null
    }
  }

  meleeHit (weapon) {
    if (!this.mesh) {
      return
    }
    return new MeleeHitArea(
      weapon.Properties.hit_radius,
      weapon.Properties.hit_angle,
      weapon.Properties.cooldown,
      this.meshRotation,
      this.mesh.position,
      this.scene
    )
  }

  playAnimation (name, loop = true) {
    if (this.container && this.currentAnimation !== name) {
      if (!this.container.animationGroups.some(ag => ag.name.includes(` ${name}`))) {
        return
      }
      this.container.animationGroups.forEach((ag) => {
        if (ag.name.includes(` ${name}`)) {
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

export default MobObserver
