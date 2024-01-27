import { EventBus } from "~/plugins/game/event_bus";
import Atlas from "~/plugins/game/atlas/atlas";
import GameObserver from "~/plugins/game/game_observer";
import freezeMaterials from "~/plugins/game/utils/freeze_materials";

const LIFTED_Y = 3

class ItemObserver {
  constructor(state) {
    this.scene = null;
    this.state = state;
    this.container = null;
    this.mesh = null;
    this.meshRotation = 0
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
    const mesh = Atlas.get(this.state.kind + "Item").clone("item-" + this.state.id);
    mesh.setParent(null)
    freezeMaterials(mesh, this.scene)
    mesh.name = "item-" + this.state.id;
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    if (this.state.rotation) {
      const rotationDelta = this.meshRotation - this.state.rotation;
      if (rotationDelta != 0) {
        this.meshRotation = this.state.rotation;
        mesh.rotate(BABYLON.Axis.Y, rotationDelta);
      }
    }
    mesh.metadata = {
      x: this.state.x,
      y: this.state.y,
      id: this.state.id,
      state: this.state
    };
    mesh.setEnabled(true)
    mesh.doNotSyncBoundingInfo = true
    this.mesh = mesh
    this.updateLifted()
  }

  updateLifted() {
    if (this.state.liftedBy !== undefined && this.state.liftedBy !== null) {
      this.mesh.unfreezeWorldMatrix()
      this.mesh.position.x = this.state.x
      this.mesh.position.z = this.state.y
      this.mesh.position.y = LIFTED_Y
    } else {
      this.mesh.position.x = this.state.x
      this.mesh.position.z = this.state.y
      this.mesh.position.y = 0
      const rotationDelta = this.meshRotation - this.state.rotation;
      if (rotationDelta != 0) {
        this.meshRotation = this.state.rotation;
        this.mesh.rotate(BABYLON.Axis.Y, rotationDelta);
      }
      this.mesh.freezeWorldMatrix()
    }
  }

  remove() {
    EventBus.$off("scene-created", this.sceneCreatedCallback);
    this.mesh.dispose();
    this.mesh = null;
    this.state = null;
  }
}

export default ItemObserver;
