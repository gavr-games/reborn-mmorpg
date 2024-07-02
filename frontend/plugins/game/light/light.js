import * as BABYLON from 'babylonjs'

class Light {
  constructor (scene) {
    this.scene = scene
    this.light = null
    this.skybox = null
  }

  create () {
    // Skybox
    this.skybox = BABYLON.MeshBuilder.CreateBox('skyBox', { size: 5000.0 }, this.scene)
    const skyboxMaterial = new BABYLON.StandardMaterial('skyBox', this.scene)
    skyboxMaterial.backFaceCulling = false
    skyboxMaterial.reflectionTexture = new BABYLON.CubeTexture('/game_assets/textures/TropicalSunnyDay', this.scene)
    skyboxMaterial.reflectionTexture.coordinatesMode = BABYLON.Texture.SKYBOX_MODE
    skyboxMaterial.diffuseColor = new BABYLON.Color3(0, 0, 0)
    skyboxMaterial.specularColor = new BABYLON.Color3(0, 0, 0)
    skyboxMaterial.disableLighting = true
    this.skybox.material = skyboxMaterial

    this.light = new BABYLON.HemisphericLight(
      'HemiLight',
      new BABYLON.Vector3(0, 1, 0),
      this.scene
    )
    this.light.intensity = 1.3
  }
}

export default Light
