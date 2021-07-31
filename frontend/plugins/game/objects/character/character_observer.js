import Atlas from "~/plugins/game/atlas/atlas";
import Camera from "~/plugins/game/camera/camera";
import { EventBus } from "~/plugins/game/event_bus";

class Character {
  constructor(state) {
    this.scene = null
    this.canvas = null
    this.state = state
    this.container = null
    this.mesh = null
    this.camera = null
    this.sceneCreatedCallback = (scene, canvas) => {
      this.scene = scene;
      this.canvas = canvas;
      this.create();
    };
    EventBus.$on("scene-created", this.sceneCreatedCallback);
  }

  create() {
    this.container = Atlas.get("baseCharacter").instantiateModelsToScene();
    this.playAnimation("Idle");
    this.mesh = this.container.rootNodes[0];
    this.mesh.setEnabled(true);
    this.mesh.position.x = this.state.x;
    this.mesh.position.z = this.state.y;
    this.camera = new Camera(this.scene, this.canvas, this);
    this.camera.create();
  }

  playAnimation(name, loop = true) {
    if (this.container) {
      this.container.animationGroups.forEach(ag => {
        if (ag.name === name) {
          ag.start(loop);
          this.currentAnimation = name;
        } else {
          ag.reset();
          ag.stop();
        }
      });
    }
  }
}

export default Character;
