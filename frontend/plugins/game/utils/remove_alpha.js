const ALPHA = 0.2

export default function removeAlpha (mesh, scene, alpha = ALPHA) {
  if (mesh.material && mesh.metadata && mesh.material.alpha === alpha) {
    mesh.material = mesh.metadata.oldMaterial
  }
  mesh.getChildren().forEach((child) => {
    removeAlpha(child, scene, alpha)
  })
}
