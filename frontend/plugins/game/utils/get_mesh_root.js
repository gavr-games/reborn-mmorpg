export default function getMeshRoot (mesh) {
  if (mesh.metadata && mesh.metadata.id) {
    return mesh
  }

  if (mesh.parent) {
    return getMeshRoot(mesh.parent)
  } else {
    return null
  }
}
