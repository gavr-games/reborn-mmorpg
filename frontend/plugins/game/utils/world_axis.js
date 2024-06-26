import * as BABYLON from 'babylonjs'

function makeTextPlane (text, color, size, scene) {
  const dynamicTexture = new BABYLON.DynamicTexture(
    'DynamicTexture',
    50,
    scene,
    true
  )
  dynamicTexture.hasAlpha = true
  dynamicTexture.drawText(
    text,
    5,
    40,
    'bold 36px Arial',
    color,
    'transparent',
    true
  )
  const plane = BABYLON.Mesh.CreatePlane('TextPlane', size, scene, true)
  plane.material = new BABYLON.StandardMaterial(
    'TextPlaneMaterial',
    scene
  )
  plane.material.backFaceCulling = false
  plane.material.specularColor = new BABYLON.Color3(0, 0, 0)
  plane.material.diffuseTexture = dynamicTexture
  return plane
}

export default function showWorldAxis (size, scene) {
  const axisX = BABYLON.Mesh.CreateLines(
    'axisX',
    [
      BABYLON.Vector3.Zero(),
      new BABYLON.Vector3(size, 0, 0),
      new BABYLON.Vector3(size * 0.95, 0.05 * size, 0),
      new BABYLON.Vector3(size, 0, 0),
      new BABYLON.Vector3(size * 0.95, -0.05 * size, 0)
    ],
    scene
  )
  axisX.color = new BABYLON.Color3(1, 0, 0)
  const xChar = makeTextPlane('X', 'red', size / 10, scene)
  xChar.position = new BABYLON.Vector3(0.9 * size, -0.05 * size, 0)
  const axisY = BABYLON.Mesh.CreateLines(
    'axisY',
    [
      BABYLON.Vector3.Zero(),
      new BABYLON.Vector3(0, size, 0),
      new BABYLON.Vector3(-0.05 * size, size * 0.95, 0),
      new BABYLON.Vector3(0, size, 0),
      new BABYLON.Vector3(0.05 * size, size * 0.95, 0)
    ],
    scene
  )
  axisY.color = new BABYLON.Color3(0, 1, 0)
  const yChar = makeTextPlane('Y', 'green', size / 10, scene)
  yChar.position = new BABYLON.Vector3(0, 0.9 * size, -0.05 * size)
  const axisZ = BABYLON.Mesh.CreateLines(
    'axisZ',
    [
      BABYLON.Vector3.Zero(),
      new BABYLON.Vector3(0, 0, size),
      new BABYLON.Vector3(0, -0.05 * size, size * 0.95),
      new BABYLON.Vector3(0, 0, size),
      new BABYLON.Vector3(0, 0.05 * size, size * 0.95)
    ],
    scene
  )
  axisZ.color = new BABYLON.Color3(0, 0, 1)
  const zChar = makeTextPlane('Z', 'blue', size / 10, scene)
  zChar.position = new BABYLON.Vector3(0, 0.05 * size, 0.9 * size)
}
