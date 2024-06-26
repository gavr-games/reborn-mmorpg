import * as BABYLON from 'babylonjs'
import Atlas from '~/plugins/game/atlas/atlas'
import Camera from '~/plugins/game/camera/camera'
import GameObserver from '~/plugins/game/game_observer'
import HealthBar from '~/plugins/game/components/health_bar'
import Nickname from '~/plugins/game/components/nickname'
import MeleeHitArea from '~/plugins/game/components/melee_hit_area'
import HighlightShape from '~/plugins/game/components/highlight_shape'
import freezeMaterials from '~/plugins/game/utils/freeze_materials'
import { EventBus } from '~/plugins/game/event_bus'

const POSITION_FIX_TRESHOLD = 1.0 // if we have difference like this between engine and client we need to urgently fix it
const ALLOWED_POSITION_DELTA = 0.1 // sync stoped object position between engine and client with this accuracy

class Character {
  constructor (state, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.scene = null
    this.canvas = null
    this.state = state
    this.container = null
    this.modelName = 'base'
    this.mesh = null
    this.meshRotation = Math.PI / 2
    this.camera = null
    this.currentAnimation = null
    this.healthbar = null
    this.targetHighlight = null
    this.dst = null
    if (state.payload.Properties.slots.body !== null) { // check if some kind of armor equipped, which changes char's look
      this.modelName = state.payload.Properties.slots.body.kind
    }
    this.pickupCallback = (params) => {
      if (params.character_id === this.state.id) {
        this.playAnimation('PickUp', false)
      }
    }
    this.startActionCallback = (params) => {
      if (params.object.player_id === this.state.player_id) {
        this.playAnimation('Punch')
      }
    }
    this.cancelActionCallback = (params) => {
      if (params.object.player_id === this.state.player_id) {
        this.playAnimation('Idle')
      }
    }
    this.sceneCreatedCallback = (scene, canvas) => {
      this.scene = scene
      this.canvas = canvas
      this.create()
      EventBus.$on('pickup_object', this.pickupCallback)
      EventBus.$on('start_delayed_action', this.startActionCallback)
      EventBus.$on('cancel_delayed_action', this.cancelActionCallback)
      EventBus.$on('finish_delayed_action', this.cancelActionCallback)
    }
    if (GameObserver.loaded) {
      this.sceneCreatedCallback(GameObserver.scene, GameObserver.canvas)
    } else {
      EventBus.$on('scene-created', this.sceneCreatedCallback)
    }
  }

  create (addRenderObserver = true) {
    this.container = Atlas.get(`${this.modelName}Character`).instantiateModelsToScene()
    this.playAnimation('Idle')
    this.mesh = this.container.rootNodes[0]
    freezeMaterials(this.mesh, this.scene)
    this.mesh.setParent(null)
    this.mesh.setEnabled(true)
    this.mesh.position.x = this.state.x
    this.mesh.position.z = this.state.y
    if (this.state.rotation) {
      const rotationDelta = this.meshRotation - this.state.rotation
      if (rotationDelta !== 0) {
        this.meshRotation = this.state.rotation
        this.mesh.rotate(BABYLON.Axis.Y, rotationDelta)
      }
    }
    // character of the logged in player
    if (this.state.player_id === this.myCharacterId && this.camera === null) {
      this.camera = new Camera(this.scene, this.canvas, this)
      this.camera.create()
    }
    this.healthbar = new HealthBar(this.state.health, this.state.max_health, this.mesh.position, this.scene)
    this.nickname = new Nickname(this.state.name, this.scene)
    this.updateAdditionalObjects()
    if (addRenderObserver) {
      GameObserver.addRenderObserver(`character-${this.state.id}`, this)
    }
  }

  update (renderInterval) {
    if (this.mesh === null) {
      return
    }
    if (this.state.speed_x !== 0 || this.state.speed_y !== 0) {
      this.playAnimation('Walk')
      this.mesh.position.x = this.mesh.position.x + this.state.speed_x / 1000 * renderInterval
      this.mesh.position.z = this.mesh.position.z + this.state.speed_y / 1000 * renderInterval
      const rotationAngle = Math.atan2(
        this.state.speed_y,
        this.state.speed_x
      )
      const rotationDelta = this.meshRotation - rotationAngle
      if (rotationDelta !== 0) {
        this.meshRotation = rotationAngle
        this.mesh.rotate(BABYLON.Axis.Y, rotationDelta)
      }
      this.updateAdditionalObjects()
    } else if (this.currentAnimation !== 'PickUp' && this.currentAnimation !== 'Punch') {
      if (this.dst === null) {
        this.playAnimation('Idle')
      } else {
        // we need to correct difference between client and engine after object stopped moving
        // move object to the correct coord
        if (Math.abs(this.mesh.position.x - this.state.x) < ALLOWED_POSITION_DELTA && Math.abs(this.mesh.position.z - this.state.y) < ALLOWED_POSITION_DELTA) {
          this.dst = null
          return
        }
        const dx = this.dst.x - this.mesh.position.x
        const dy = this.dst.y - this.mesh.position.z
        const angle = Math.atan2(dy, dx)
        const speedX = this.state.speed * Math.cos(angle)
        const speedY = this.state.speed * Math.sin(angle)
        this.mesh.position.x = this.mesh.position.x + speedX / 1000 * renderInterval
        this.mesh.position.z = this.mesh.position.z + speedY / 1000 * renderInterval
        this.updateAdditionalObjects()
      }
    }
  }

  updateAdditionalObjects () {
    if (this.mesh === null) {
      return
    }
    if (this.state.liftedObjectId !== undefined && this.state.liftedObjectId !== null) {
      const liftedMesh = this.scene.getNodeByName(`item-${this.state.liftedObjectId}`)
      if (liftedMesh !== null) {
        liftedMesh.position.x = this.mesh.position.x
        liftedMesh.position.z = this.mesh.position.y
      }
    }
    this.healthbar.update(this.state.health, this.state.max_health, this.mesh.position)
    this.nickname.update(this.mesh.position)
    if (this.targetHighlight) {
      this.targetHighlight.update(this.mesh.position)
    }
    // character of the logged in player
    if (this.state.player_id === this.myCharacterId) {
      this.camera.update(this.mesh.position)
    }
  }

  updateFromEngine () {
    if (!this.mesh) {
      return
    }
    this.dst = null
    // Character stopped, but we have difference in positions between engine and client
    if (this.state.speed_x === 0 && this.state.speed_y === 0 && (this.mesh.position.x !== this.state.x || this.mesh.position.y !== this.state.y)) {
      this.dst = { x: this.state.x, y: this.state.y }
    } else if (Math.abs(this.mesh.position.x - this.state.x) > POSITION_FIX_TRESHOLD || Math.abs(this.mesh.position.z - this.state.y) > POSITION_FIX_TRESHOLD) {
      // the difference between client and engine positions for moving object is too big, teleport object
      this.mesh.position.x = this.state.x
      this.mesh.position.z = this.state.y
    }
  }

  remove () {
    GameObserver.removeRenderObserver(`character-${this.state.id}`)
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    EventBus.$off('pickup_object', this.pickupCallback)
    EventBus.$off('start_delayed_action', this.startActionCallback)
    EventBus.$off('cancel_delayed_action', this.cancelActionCallback)
    EventBus.$off('finish_delayed_action', this.cancelActionCallback)
    if (this.state.player_id === this.myCharacterId && this.camera !== null) {
      this.camera.remove()
      this.camera = null
    }
    this.state = null
    if (this.container !== null) {
      this.container.dispose()
    }
    this.container = null
    this.removeMesh()
  }

  removeMesh () {
    if (this.healthbar) {
      this.healthbar.remove()
      this.healthbar = null
    }
    if (this.nickname) {
      this.nickname.remove()
      this.nickname = null
    }
    if (this.targetHighlight) {
      this.targetHighlight.remove()
      this.targetHighlight = null
    }
    if (this.mesh !== null) {
      this.mesh.dispose()
      this.mesh = null
    }
  }

  changeModel (model) {
    this.modelName = model
    this.meshRotation = Math.PI / 2
    this.removeMesh()
    this.create(false)
    if (this.state.player_id === this.myCharacterId && this.camera !== null) {
      this.camera.updateLockedTarget(this.mesh)
    }
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

export default Character
