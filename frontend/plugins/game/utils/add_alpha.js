import * as BABYLON from 'babylonjs'
const ALPHA = 0.2

export default function addAlpha (mesh, scene, alpha = ALPHA) {
  if (mesh.material && mesh.material.alpha !== alpha) {
    const alphaMaterial = new BABYLON.StandardMaterial(`alphaMaterial-${mesh.metadata.id}`, scene)

    alphaMaterial.diffuseColor = new BABYLON.Color3(1, 0, 1)
    alphaMaterial.specularColor = new BABYLON.Color3(0.5, 0.6, 0.87)
    alphaMaterial.emissiveColor = new BABYLON.Color3(1, 1, 1)
    alphaMaterial.ambientColor = new BABYLON.Color3(0.23, 0.98, 0.53)
    alphaMaterial.alpha = ALPHA
    mesh.metadata.oldMaterial = mesh.material
    mesh.material = alphaMaterial
  }
  mesh.getChildren().forEach((child) => {
    addAlpha(child, scene)
  })
}
