import * as BABYLON from 'babylonjs'
import freezeMaterials from '~/plugins/game/utils/freeze_materials'

/**
 * Creates a new ContainerAssetTask
 * @param name defines the name of the task
 * @param meshesNames defines the list of mesh's names you want to load
 * @param rootUrl defines the root url to use as a base to load your meshes and associated resources
 * @param sceneFilename defines the filename of the scene to load from
 */
class ContainerAssetTask extends BABYLON.AbstractAssetTask {
  constructor (
    /**
     * Defines the name of the task
     */
    name,
    /**
     * Defines the list of mesh's names you want to load
     */
    meshesNames,
    /**
     * Defines the root url to use as a base to load your meshes and associated resources
     */
    rootUrl,
    /**
     * Defines the filename of the scene to load from
     */
    sceneFilename
  ) {
    super(name)
    this.name = name
    this.meshesNames = meshesNames
    this.rootUrl = rootUrl
    this.sceneFilename = sceneFilename
  }

  /**
   * Execute the current task
   * @param scene defines the scene where you want your assets to be loaded
   * @param onSuccess is a callback called when the task is successfully executed
   * @param onError is a callback called if an error occurs
   */
  runTask (scene, onSuccess, onError) {
    const _this = this
    BABYLON.SceneLoader.LoadAssetContainer(
      this.rootUrl,
      this.sceneFilename,
      scene,
      function (container) {
        _this.loadedContainer = container
        _this.loadedMeshes = container.meshes
        _this.loadedParticleSystems = container.particleSystems
        _this.loadedSkeletons = container.skeletons
        _this.loadedAnimationGroups = container.animationGroups
        onSuccess()
      },
      null,
      function (scene, message, exception) {
        onError(message, exception)
      }
    )
  }
}
BABYLON.AssetsManager.prototype.addContainerTask = function (
  taskName,
  meshesNames,
  rootUrl,
  sceneFilename
) {
  const task = new ContainerAssetTask(
    taskName,
    meshesNames,
    rootUrl,
    sceneFilename
  )
  this._tasks.push(task)
  return task
}

class Loader {
  constructor (scene, finishCallback, atlas) {
    this.atlas = atlas
    this.scene = scene
    this.finishCallback = finishCallback
    this.assetsManager = new BABYLON.AssetsManager(scene)
    this.assetsManager.addContainerTask = function (
      taskName,
      meshesNames,
      rootUrl,
      sceneFilename
    ) {
      const task = new ContainerAssetTask(
        taskName,
        meshesNames,
        rootUrl,
        sceneFilename
      )
      this._tasks.push(task)
      return task
    }
    this.assetsManager.onFinish = finishCallback
  }

  taskToMesh (task) {
    const mesh = task.loadedContainer.instantiateModelsToScene().rootNodes[0].getChildren()[0]
    mesh.setEnabled(false)
    freezeMaterials(mesh, this.scene)
    return mesh
  }

  load () {
    this.loadCharacters()
    this.loadSurfaces()
    this.loadRocks()
    this.loadTrees()
    this.loadPlants()
    this.loadItems()
    this.loadMobs()
    this.loadNpcs()
    this.assetsManager.load()
  }

  loadCharacters () {
    const characters = [
      'base',
      'golden_armor',
      'leather_robe'
    ]
    characters.forEach((character) => {
      const task = this.assetsManager.addContainerTask(
        character,
        character,
        '/game_assets/characters/',
        character + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(character + 'Character', task.loadedContainer)
      }
    })
  }

  loadSurfaces () {
    const surfaces = [
      'grass',
      'dirt',
      'water',
      'sand',
      'stone'
    ]
    surfaces.forEach((surface) => {
      const task = this.assetsManager.addContainerTask(
        surface,
        surface,
        '/game_assets/surfaces/',
        surface + '.glb'
      )
      task.onSuccess = (task) => {
        // This should be changed if surface models of aother structure are imported
        this.atlas.set(surface + 'Surface', this.taskToMesh(task))
      }
    })
  }

  loadRocks () {
    const rocks = [
      'rock_moss'
    ]
    rocks.forEach((rock) => {
      const task = this.assetsManager.addContainerTask(
        rock,
        rock,
        '/game_assets/rocks/',
        rock + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(rock + 'Rock', this.taskToMesh(task))
      }
    })
  }

  loadTrees () {
    const trees = [
      'tree_5',
      'pine_5'
    ]
    trees.forEach((tree) => {
      const task = this.assetsManager.addContainerTask(
        tree,
        tree,
        '/game_assets/trees/',
        tree + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(tree + 'Tree', this.taskToMesh(task))
      }
    })
  }

  loadPlants () {
    const plants = [
      'cactus',
      'grass_plant',
      'carrot_sprout',
      'carrot_ripe'
    ]
    plants.forEach((plant) => {
      const task = this.assetsManager.addContainerTask(
        plant,
        plant,
        '/game_assets/plants/',
        plant + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(plant + 'Plant', this.taskToMesh(task))
      }
    })
  }

  loadItems () {
    const items = [
      'anvil',
      'axe',
      'animal_skin',
      'bone',
      'bone_needle',
      'bone_saw',
      'carrot',
      'carrot_seed',
      'claim_stone',
      'claim_obelisk',
      'cactus_slice',
      'fire_dragon_egg',
      'fire_dragon_hatchery',
      'gold',
      'gold_ingot',
      'golden_armor',
      'grass',
      'healing_balm',
      'iron_ingot',
      'iron_nails',
      'leather_robe',
      'log',
      'pickaxe',
      'rope',
      'small_bag',
      'stone',
      'stone_hammer',
      'stone_spear',
      'stone_knife',
      'stone_wall',
      'wooden_chest',
      'wooden_fishing_rod',
      'wooden_shovel',
      'wooden_wall'
    ]
    items.forEach((item) => {
      const task = this.assetsManager.addContainerTask(
        item,
        item,
        '/game_assets/items/',
        item + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(item + 'Item', this.taskToMesh(task))
      }
    })
  }

  loadMobs () {
    const mobs = [
      'fire_dragon',
      'bat'
    ]
    mobs.forEach((item) => {
      const task = this.assetsManager.addContainerTask(
        item,
        item,
        '/game_assets/mobs/',
        item + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(item + 'Mob', task.loadedContainer)
      }
    })
  }

  loadNpcs () {
    const npcs = [
      'town_keeper'
    ]
    npcs.forEach((item) => {
      const task = this.assetsManager.addContainerTask(
        item,
        item,
        '/game_assets/npcs/',
        item + '.glb'
      )
      task.onSuccess = (task) => {
        this.atlas.set(item + 'Npc', task.loadedContainer)
      }
    })
  }
}

export default Loader
