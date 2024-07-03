import * as BABYLON from 'babylonjs'
import { FurMaterial, WaterMaterial } from 'babylonjs-materials'
import GameObserver from '~/plugins/game/game_observer'
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
      'dungeon_floor',
      'water',
      'sand',
      'stone',
      'stone_road',
      'town_floor'
    ]
    surfaces.forEach((surface) => {
      const task = this.assetsManager.addContainerTask(
        surface,
        surface,
        '/game_assets/surfaces/',
        surface + '.glb'
      )
      task.onSuccess = (task) => {
        let mesh = this.taskToMesh(task)
        if (surface === 'grass') {
          // Ground template
          const grassMaterial = new FurMaterial('grass-material', this.scene)
          grassMaterial.highLevelFur = false
          grassMaterial.furLength = 0.3 // Represents the maximum length of the fur, which is then adjusted randomly. Default value is 1.
          grassMaterial.furAngle = 0 // Represents the angle the fur lies on the mesh from 0 to Math.PI/2. The default angle of 0 gives fur sticking straight up and PI/2 lies along the mesh.
          grassMaterial.furColor = new BABYLON.Color3(0.02, 0.61, 0.46)
          const grass = BABYLON.MeshBuilder.CreateGround('grass', { height: 1, width: 1, subdivisions: 50 }, this.scene)
          grass.position.y = 0
          grass.material = grassMaterial
          grass.position.x = -100
          grass.position.z = -100
          grass.doNotSyncBoundingInfo = true
          grass.isPickable = false
          grass.freezeWorldMatrix()
          grass.material.freeze()
          mesh = grass
        }
        if (surface === 'water') {
          const water = new WaterMaterial('water', this.scene, new BABYLON.Vector2(512, 512))
          water.backFaceCulling = true
          water.bumpTexture = new BABYLON.Texture('/game_assets/textures/waterbump.png', this.scene)
          water.windForce = 0
          water.waveHeight = 0.0
          water.bumpHeight = 0.05
          water.windDirection = new BABYLON.Vector2(1, 1)
          water.waterColor = new BABYLON.Color3(0, 0, 221 / 255)
          water.colorBlendFactor = 0.0
          water.addToRenderList(GameObserver.light.skybox)
          mesh.material = water
        }
        this.atlas.set(surface + 'Surface', mesh)
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
      'palm_3',
      'pine_5',
      'tree_5'
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
      'carrot_ripe',
      'tomato_sprout',
      'tomato_ripe'
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
      'bell_tower',
      'blacksmith',
      'bone',
      'bone_needle',
      'bone_saw',
      'bonfire',
      'bonfire_burning',
      'bonfire_extinguished',
      'brick_column',
      'brick_wall',
      'carrot',
      'carrot_seed',
      'claim_stone',
      'claim_obelisk',
      'cactus_slice',
      'dragon_altar',
      'dungeon_chest',
      'dungeon_column',
      'dungeon_exit',
      'dungeon_key',
      'dungeon_wall',
      'fire_dragon_egg',
      'fire_dragon_hatchery',
      'fish',
      'fried_fish',
      'iron_frying_pan',
      'gold',
      'gold_ingot',
      'golden_armor',
      'grass',
      'healing_balm',
      'house',
      'inn',
      'iron_ingot',
      'iron_nails',
      'leather_robe',
      'log',
      'market_stand',
      'pickaxe',
      'rope',
      'sawmill',
      'small_bag',
      'stone',
      'stone_hammer',
      'stone_spear',
      'stone_knife',
      'stone_wall',
      'tomato',
      'tomato_seed',
      'town_gate',
      'trapdoor',
      'well',
      'windmill',
      'wooden_bench',
      'wooden_chest',
      'wooden_door',
      'wooden_door_closed',
      'wooden_door_opened',
      'wooden_fence',
      'wooden_fishing_rod',
      'wooden_shovel',
      'wooden_table',
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
      'bat',
      'baby_fire_dragon',
      'fire_dragon',
      'zombie'
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
      'town_keeper',
      'dungeon_keeper'
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
