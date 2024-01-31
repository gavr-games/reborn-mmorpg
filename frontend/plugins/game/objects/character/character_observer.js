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

class Character {
  constructor (state, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.scene = null
    this.canvas = null
    this.state = state
    this.container = null
    this.mesh = null
    this.meshRotation = Math.PI / 2
    this.camera = null
    this.currentAnimation = null
    this.healthbar = null
    this.targetHighlight = null
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

  create () {
    this.container = Atlas.get('baseCharacter').instantiateModelsToScene()
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
    if (this.state.player_id === this.myCharacterId) {
      this.camera = new Camera(this.scene, this.canvas, this)
      this.camera.create()
    }
    this.healthbar = new HealthBar(this.state.health, this.state.max_health, this.mesh.position, this.scene)
    this.nickname = new Nickname(this.state.name, this.scene)
    GameObserver.addRenderObserver(`character-${this.state.id}`, this)
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
    } else if (this.currentAnimation !== 'PickUp' && this.currentAnimation !== 'Punch') {
      this.playAnimation('Idle')
    }
    this.mesh.position.x = this.state.x
    this.mesh.position.z = this.state.y
    if (this.state.liftedObjectId !== undefined && this.state.liftedObjectId !== null) {
      const liftedMesh = this.scene.getNodeByName(`item-${this.state.liftedObjectId}`)
      if (liftedMesh !== null) {
        liftedMesh.position.x = this.state.x
        liftedMesh.position.z = this.state.y
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

  remove () {
    GameObserver.removeRenderObserver(`character-${this.state.id}`)
    EventBus.$off('scene-created', this.sceneCreatedCallback)
    EventBus.$off('pickup_object', this.pickupCallback)
    EventBus.$off('start_delayed_action', this.startActionCallback)
    EventBus.$off('cancel_delayed_action', this.cancelActionCallback)
    EventBus.$off('finish_delayed_action', this.cancelActionCallback)
    this.healthbar.remove()
    this.healthbar = null
    this.nickname.remove()
    this.nickname = null
    if (this.targetHighlight) {
      this.targetHighlight.remove()
      this.targetHighlight = null
    }
    this.mesh.dispose()
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

export default Character
