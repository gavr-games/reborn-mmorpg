import * as BABYLON from 'babylonjs'
import { AdvancedDynamicTexture, TextBlock } from 'babylonjs-gui'

const TEXT_SIZE = 24
const TEXT_Y = 4

class Nickname {
  constructor (nickname, scene) {
    this.scene = scene
    this.nickname = nickname
    this.plane = BABYLON.MeshBuilder.CreatePlane('nickname-plane', { height: 20, width: 20 }, scene)
    this.plane.position.y = TEXT_Y
    this.plane.billboardMode = BABYLON.Mesh.BILLBOARDMODE_ALL
    this.advancedTexture = AdvancedDynamicTexture.CreateForMesh(this.plane)
    this.text = new TextBlock()
    this.text.text = nickname
    this.text.color = 'black'
    this.text.fontSize = TEXT_SIZE
    this.advancedTexture.addControl(this.text)
    this.plane.convertToUnIndexedMesh()
    this.plane.isPickable = false
  }

  update (pos) {
    if (this.plane) {
      this.plane.position.x = pos.x
      this.plane.position.z = pos.z
    }
  }

  remove () {
    if (this.plane) {
      this.plane.dispose()
    }
    this.plane = null
    this.advancedTexture = null
    this.text = null
  }
}

export default Nickname
