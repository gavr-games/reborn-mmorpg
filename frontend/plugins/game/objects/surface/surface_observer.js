import { EventBus } from "~/plugins/game/event_bus";
import Atlas from "~/plugins/game/atlas/atlas";
import GameObserver from "~/plugins/game/game_observer";

class SurfaceObserver {
  constructor(state) {
    this.scene = null;
    this.state = state;
    this.container = null;
    this.mesh = null;
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
    let mesh = Atlas.get(this.state.kind + "Surface").createInstance("surface-" + this.state.id)
    mesh.setParent(null)
    mesh.isPickable = false
    mesh.name = "surface-" + this.state.id
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    if (this.state.rotation) {
      mesh.rotate(BABYLON.Axis.Y, Math.PI / 2);
    }
    mesh.metadata = {
      x: this.state.x,
      y: this.state.y,
      id: this.state.id,
      state: this.state
    };
    mesh.setEnabled(true);
    mesh.freezeWorldMatrix();
    mesh.doNotSyncBoundingInfo = true;
    this.mesh = mesh;
  }

  remove() {
    EventBus.$off("scene-created", this.sceneCreatedCallback);
    if (this.mesh) {
      this.mesh.dispose();
    }
    this.mesh = null;
    this.state = null;
  }
}

export default SurfaceObserver;
