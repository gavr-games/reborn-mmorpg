import { EventBus } from "~/plugins/game/event_bus";
import Atlas from "~/plugins/game/atlas/atlas";
import GameObserver from "~/plugins/game/game_observer";
import addAlpha from "~/plugins/game/utils/add_alpha";

class CraftObserver {
  constructor() {
    this.scene = null;
    this.container = null;
    this.mesh = null;
    this.rotation = 0
    this.plane = null // this plane is used for picking coords on pointer moved
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene;
    } else {
      EventBus.$on("scene-created", scene => {
        this.scene = scene;
      });
    }
  }

  create(itemKey, x, y) {
    let mesh = Atlas.get(itemKey + "Item").clone("craft-item");
    mesh.setParent(null)
    mesh.setEnabled(true);
    mesh.name = "craft-item"
    mesh.position.x = x
    mesh.position.y = 0
    mesh.position.z = y
    mesh.setEnabled(true);
    mesh.doNotSyncBoundingInfo = true;
    mesh.isPickable = false
    this.mesh = mesh;
    addAlpha(this.mesh, this.scene, 0.7)
    GameObserver.grid.create()
    // replace 1000 with floor width
    this.plane = BABYLON.MeshBuilder.CreatePlane("craft-plane", {height: 1000, width: 1000}, this.scene)
    this.plane.position.x = 0
    this.plane.position.z = 0
    this.plane.position.y = 0.05
    this.plane.rotate(BABYLON.Axis.X, Math.PI / 2);
    this.plane.material = new BABYLON.StandardMaterial("ObjectplaneMaterial", this.scene)
    this.plane.material.alpha = 0;
    this.plane.convertToUnIndexedMesh()
    this.plane.doNotSyncBoundingInfo = true
  }

  update(x, y) {
    this.mesh.position.x = x
    this.mesh.position.z = y
  }

  rotate() {
    let angle = Math.PI / 2
    if (this.rotation == 1) {
      angle = - Math.PI / 2
    }
    this.rotation = this.rotation == 0 ? 1 : 0
    this.mesh.rotate(BABYLON.Axis.Y, angle);
  }

  remove() {
    EventBus.$off("scene-created", this.sceneCreatedCallback);
    if (this.mesh) {
      this.mesh.dispose();
      this.mesh = null;

    }
    this.plane.dispose()
    this.plane = null
    GameObserver.grid.remove()
  }
}

export default CraftObserver;
