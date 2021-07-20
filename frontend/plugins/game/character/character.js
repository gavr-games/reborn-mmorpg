import * as BABYLON from "babylonjs";
import Atlas from "~/plugins/game/atlas/atlas";

class Character {
  constructor(scene, canvas) {
    this.scene = scene;
    this.canvas = canvas;
    this.container = null
    this.mesh = null
  }

  create() {
    this.container = Atlas.get("baseCharacter").instantiateModelsToScene();
    this.playAnimation("Walk");
    this.mesh = this.container.rootNodes[0];
    this.mesh.setEnabled(true);
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
