import Atlas from "~/plugins/game/atlas/atlas";
import Camera from "~/plugins/game/camera/camera";
import GameObserver from "~/plugins/game/game_observer";
import { EventBus } from "~/plugins/game/event_bus";

class Character {
  constructor(state, myCharacterId) {
    this.myCharacterId = myCharacterId
    this.scene = null
    this.canvas = null
    this.state = state
    this.container = null
    this.mesh = null
    this.meshRotation = Math.PI / 2
    this.camera = null
    this.currentAnimation = null
    this.pickupCallback = (params) => {
      if (params.character_id === this.state.id) {
        this.playAnimation("PickUp", false)
      }
    }
    this.startActionCallback = (params) => {
      if (params.object.player_id === this.state.player_id) {
        this.playAnimation("Punch")
      }
    }
    this.cancelActionCallback = (params) => {
      if (params.object.player_id === this.state.player_id) {
        this.playAnimation("Idle")
      }
    }
    this.sceneCreatedCallback = (scene, canvas) => {
      this.scene = scene
      this.canvas = canvas
      this.create()
      EventBus.$on("pickup_object", this.pickupCallback)
      EventBus.$on("start_delayed_action", this.startActionCallback)
      EventBus.$on("cancel_delayed_action", this.cancelActionCallback)
      EventBus.$on("finish_delayed_action", this.cancelActionCallback)
    }
    if (GameObserver.loaded) {
      this.sceneCreatedCallback(GameObserver.scene, GameObserver.canvas)
    } else {
      EventBus.$on("scene-created", this.sceneCreatedCallback)
    }
  }

  create() {
    this.container = Atlas.get("baseCharacter").instantiateModelsToScene();
    this.playAnimation("Idle");
    this.mesh = this.container.rootNodes[0];
    this.mesh.setParent(null)
    this.mesh.setEnabled(true);
    this.mesh.position.x = this.state.x;
    this.mesh.position.z = this.state.y;
    if (this.state.rotation) {
      mesh.rotate(BABYLON.Axis.Y, Math.PI / 2);
    }
    // character of the logged in player
    if (this.state.player_id == this.myCharacterId) {
      this.camera = new Camera(this.scene, this.canvas, this);
      this.camera.create();
    }
    GameObserver.addRenderObserver(`character-${this.state.id}`, this);
  }

  update(renderInterval) {
    if (this.state.speed_x != 0 || this.state.speed_y != 0) {
      this.playAnimation("Walk");
      this.state.x =  this.state.x + this.state.speed_x / 1000 * renderInterval
      this.state.y =  this.state.y + this.state.speed_y / 1000 * renderInterval
      const rotationAngle = Math.atan2(
        this.state.y - this.mesh.position.z,
        this.state.x - this.mesh.position.x
      );
      let rotationDelta = this.meshRotation - rotationAngle;
      if (rotationDelta != 0) {
        this.meshRotation = rotationAngle;
        this.mesh.rotate(BABYLON.Axis.Y, rotationDelta);
      }
    } else if (this.currentAnimation != "PickUp" && this.currentAnimation != "Punch") {
      this.playAnimation("Idle");
    }
    this.mesh.position.x = this.state.x
    this.mesh.position.z = this.state.y
    // character of the logged in player
    if (this.state.player_id == this.myCharacterId) {
      this.camera.update(this.mesh.position)
    }
  }

  remove() {
    GameObserver.removeRenderObserver(`character-${this.state.id}`);
    EventBus.$off("scene-created", this.sceneCreatedCallback);
    EventBus.$off("pickup_object", this.pickupCallback)
    EventBus.$off("start_delayed_action", this.startActionCallback)
    EventBus.$off("cancel_delayed_action", this.cancelActionCallback)
    EventBus.$off("finish_delayed_action", this.cancelActionCallback)
    this.mesh.dispose();
    this.mesh = null;
    this.state = null;
  }

  playAnimation(name, loop = true) {
    if (this.container && this.currentAnimation != name) {
      this.container.animationGroups.forEach(ag => {
        if (ag.name.includes(name)) {
          ag.start(loop);
          this.currentAnimation = name;
          if (!loop) {
            ag.onAnimationEndObservable.addOnce(() => {
              this.currentAnimation = "Idle"
            });
          }
        } else {
          ag.reset();
          ag.stop();
        }
      });
    }
  }
}

export default Character;
