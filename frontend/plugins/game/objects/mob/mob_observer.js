import { EventBus } from "~/plugins/game/event_bus";
import Atlas from "~/plugins/game/atlas/atlas";
import GameObserver from "~/plugins/game/game_observer";

class MobObserver {
  constructor(state) {
    this.scene = null;
    this.state = state;
    this.container = null;
    this.mesh = null;
    this.meshRotation = Math.PI / 2
    this.currentAnimation = null
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene;
      this.create();
    } else {
      EventBus.$on("scene-created", scene => {
        this.scene = scene;
        this.create();
      });
    }

  }

  create() {
    this.container = Atlas.get(this.state.kind + "Mob").instantiateModelsToScene();
    this.playAnimation("Idle");
    let mesh = this.container.rootNodes[0];
    mesh.setParent(null)
    mesh.setEnabled(true);
    mesh.name = "mob-" + this.state.id;
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    mesh.metadata = {
      x: this.state.x,
      y: this.state.y,
      id: this.state.id,
      state: this.state
    };
    mesh.setEnabled(true);
    mesh.doNotSyncBoundingInfo = true;
    this.mesh = mesh;
    GameObserver.addRenderObserver(`mob-${this.state.id}`, this);
  }

  update(renderInterval) {
    if (this.state.speed_x != 0 || this.state.speed_y != 0) {
      //this.playAnimation("Walk");
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
    } else {
      this.playAnimation("Idle");
    }
    this.mesh.position.x = this.state.x
    this.mesh.position.z = this.state.y
  }

  remove() {
    GameObserver.removeRenderObserver(`mob-${this.state.id}`);
    EventBus.$off("scene-created", this.sceneCreatedCallback);
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

export default MobObserver;
