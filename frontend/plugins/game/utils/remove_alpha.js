const ALPHA = 0.5

export default function removeAlpha(mesh, scene) {
  if (mesh.material && mesh.material.alpha == ALPHA) {
    mesh.material = mesh.metadata.oldMaterial
  }
  mesh.getChildren().forEach((child) => {
    removeAlpha(child, scene)
  })
}
