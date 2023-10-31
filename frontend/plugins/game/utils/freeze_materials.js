export default function freezeMaterials(mesh, scene) {
  if (mesh.material) {
    mesh.material.freeze()
  }
  mesh.getChildren().forEach((child) => {
    freezeMaterials(child, scene)
  })
}
