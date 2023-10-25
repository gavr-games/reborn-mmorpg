import * as BABYLON from "babylonjs";

/**
 * Creates a new ContainerAssetTask
 * @param name defines the name of the task
 * @param meshesNames defines the list of mesh's names you want to load
 * @param rootUrl defines the root url to use as a base to load your meshes and associated resources
 * @param sceneFilename defines the filename of the scene to load from
 */
class ContainerAssetTask extends BABYLON.AbstractAssetTask {
  constructor(
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
    super(name);
    this.name = name;
    this.meshesNames = meshesNames;
    this.rootUrl = rootUrl;
    this.sceneFilename = sceneFilename;
  }
  /**
   * Execute the current task
   * @param scene defines the scene where you want your assets to be loaded
   * @param onSuccess is a callback called when the task is successfully executed
   * @param onError is a callback called if an error occurs
   */
  runTask(scene, onSuccess, onError) {
    var _this = this;
    BABYLON.SceneLoader.LoadAssetContainer(
      this.rootUrl,
      this.sceneFilename,
      scene,
      function(container) {
        _this.loadedContainer = container;
        _this.loadedMeshes = container.meshes;
        _this.loadedParticleSystems = container.particleSystems;
        _this.loadedSkeletons = container.skeletons;
        _this.loadedAnimationGroups = container.animationGroups;
        onSuccess();
      },
      null,
      function(scene, message, exception) {
        onError(message, exception);
      }
    );
  }
}
BABYLON.AssetsManager.prototype.addContainerTask = function(
  taskName,
  meshesNames,
  rootUrl,
  sceneFilename
) {
  var task = new ContainerAssetTask(
    taskName,
    meshesNames,
    rootUrl,
    sceneFilename
  );
  this._tasks.push(task);
  return task;
};

class Loader {
  constructor(scene, finishCallback, atlas) {
    this.atlas = atlas;
    this.scene = scene;
    this.finishCallback = finishCallback;
    this.assetsManager = new BABYLON.AssetsManager(scene);
    this.assetsManager.addContainerTask = function(
      taskName,
      meshesNames,
      rootUrl,
      sceneFilename
    ) {
      var task = new ContainerAssetTask(
        taskName,
        meshesNames,
        rootUrl,
        sceneFilename
      );
      this._tasks.push(task);
      return task;
    };
    this.assetsManager.onFinish = finishCallback;
  }

  load() {
    let task = this.assetsManager.addContainerTask(
      "baseCharacter",
      "baseCharacter",
      "/game_assets/characters/",
      "base_character.glb"
    )
    task.onSuccess = task => {
      this.atlas.set("baseCharacter", task.loadedContainer);
    }
    this.loadSurfaces()
    this.loadRocks()
    this.loadTrees()
    this.loadItems()
    this.loadMobs()
    this.assetsManager.load()
  }

  loadSurfaces() {
    let surfaces = [
      "grass",
      "water",
      "sand",
    ];
    surfaces.forEach(surface => {
      let task = this.assetsManager.addContainerTask(
        surface,
        surface,
        "/game_assets/surfaces/",
        surface + ".glb"
      );
      task.onSuccess = task => {
        // This should be changed if surface models of aother structure are imported
        let mesh = task.loadedContainer.instantiateModelsToScene().rootNodes[0].getChildren()[0]
        mesh.setEnabled(false)
        this.atlas.set(surface + "Surface", mesh);
      };
    });
  }

  loadRocks() {
    let rocks = [
      "rock_moss",
    ];
    rocks.forEach(rock => {
      let task = this.assetsManager.addContainerTask(
        rock,
        rock,
        "/game_assets/rocks/",
        rock + ".glb"
      );
      task.onSuccess = task => {
        this.atlas.set(rock + "Rock", task.loadedContainer);
      };
    });
  }

  loadTrees() {
    let trees = [
      "tree_5",
      "pine_5",
    ];
    trees.forEach(tree => {
      let task = this.assetsManager.addContainerTask(
        tree,
        tree,
        "/game_assets/trees/",
        tree + ".glb"
      );
      task.onSuccess = task => {
        this.atlas.set(tree + "Tree", task.loadedContainer);
      };
    });
  }

  loadItems() {
    let items = [
      "axe",
      "pickaxe",
      "log",
      "stone",
      "stone_hammer",
      "stone_wall",
      "fire_dragon_egg",
      "fire_dragon_hatchery",
      "stone_spear",
    ];
    items.forEach(item => {
      let task = this.assetsManager.addContainerTask(
        item,
        item,
        "/game_assets/items/",
        item + ".glb"
      );
      task.onSuccess = task => {
        this.atlas.set(item + "Item", task.loadedContainer);
      };
    });
  }

  loadMobs() {
    let mobs = [
      "fire_dragon",
      "bat",
    ];
    mobs.forEach(item => {
      let task = this.assetsManager.addContainerTask(
        item,
        item,
        "/game_assets/mobs/",
        item + ".glb"
      );
      task.onSuccess = task => {
        this.atlas.set(item + "Mob", task.loadedContainer);
      };
    });
  }
}

export default Loader;
