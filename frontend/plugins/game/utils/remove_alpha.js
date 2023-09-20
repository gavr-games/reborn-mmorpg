const ALPHA = 0.5

export default function removeAlpha(mesh, scene, alpha=ALPHA) {
  if (mesh.material && mesh.material.alpha == alpha) {
    mesh.material = mesh.metadata.oldMaterial
  }
  mesh.getChildren().forEach((child) => {
    removeAlpha(child, scene)
  })
}
