import * as BABYLON from "babylonjs"

const ALPHA = 0.5
const Y = 0.1

class HighlightShape {
  constructor(obj, pos, scene) {
    this.scene = scene
    if (obj.Properties["shape"]) {
      this.shape = BABYLON.MeshBuilder.CreateDisc("shape-disc", {radius: obj.Properties["width"] / 2, arc: 1, tessellation: 24}, scene)
    } else { //rectangle
      this.shape = BABYLON.MeshBuilder.CreatePlane("shape-plane", {height: obj.Properties["height"], width: obj.Properties["width"]}, scene)
    }
    this.shape.rotate(BABYLON.Axis.X, Math.PI / 2);
    this.shape.position.x = pos.x
    this.shape.position.z = pos.z
    this.shape.position.y = Y
    this.shape.material = new BABYLON.StandardMaterial("ObjectShapeMaterial", scene)
    this.shape.material.diffuseColor.set(1, 0, 0)
    this.shape.material.alpha = ALPHA;
    this.shape.convertToUnIndexedMesh()
    this.shape.doNotSyncBoundingInfo = true
    this.shape.isPickable = false
  }

  update(pos) {
    this.shape.position.x = pos.x
    this.shape.position.z = pos.z
  }

  remove() {
    if (this.shape) {
      this.shape.dispose();
    }
    this.shape = null;
  }
}

export default HighlightShape;
