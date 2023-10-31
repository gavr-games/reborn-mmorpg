import * as BABYLON from "babylonjs";
import * as Loaders from "babylonjs-loaders";
import Atlas from "~/plugins/game/atlas/atlas";
import Loader from "~/plugins/game/atlas/loader";
import Light from "~/plugins/game/light/light";
import { EventBus } from "~/plugins/game/event_bus";
import showWorldAxis from "~/plugins/game/utils/world_axis";
import Grid from "~/plugins/game/utils/grid";
import getMeshRoot from "~/plugins/game/utils/get_mesh_root";
import addAlpha from "~/plugins/game/utils/add_alpha";
import removeAlpha from "~/plugins/game/utils/remove_alpha";

const NON_TRANSPARENT_OBJECT_TYPES = ["player", "surface"]

class GameObserver {
  constructor() {
    this.canvas = null;
    this.engine = null;
    this.scene = null;
    this.loader = null;
    this.light = null;
    this.fpsEl = null;
    this.grid = null;
    this.lastTick = 0;
    this.loaded = false;
    this.renderObservers = [];
    this.previousAlphaMeshes = [];
    this.previousAlphaMeshesIds = [];
  }

  init() {
    this.canvas = document.getElementById("game-canvas")
    this.engine = new BABYLON.Engine(this.canvas, true, {
      preserveDrawingBuffer: true,
      stencil: true
    })
    Loaders.OBJFileLoader.OPTIMIZE_WITH_UV = true;
    this.createScene();
  }

  createScene() {
    this.scene = new BABYLON.Scene(this.engine, {
      useGeometryUniqueIdsMap: true,
      useMaterialMeshMap: true,
      useClonedMeshMap: true,
    });
    this.scene.actionManager = new BABYLON.ActionManager(this.scene);
    this.registerActions(this.scene);

    this.scene.useGeometryIdsMap = true
    this.scene.useMaterialMeshMap = true
    this.scene.useClonedMeshMap = true

    this.scene.performancePriority === BABYLON.ScenePerformancePriority.Aggressive

    this.loader = new Loader(
      this.scene,
      () => {
        this.createObjects();
        this.runRenderLoop();
        this.scene.registerBeforeRender(() => {
          this.castAlphaRay()
        })
        // this.scene.debugLayer.show()
        this.loaded = true;
        EventBus.$emit("scene-created", this.scene, this.canvas);
      },
      Atlas
    );
    this.loader.load();
  }

  registerActions(scene) {
    scene.actionManager.registerAction(
      new BABYLON.ExecuteCodeAction(
        BABYLON.ActionManager.OnKeyDownTrigger,
        evt => {
          EventBus.$emit("keydown", evt.sourceEvent.key);
        }
      )
    );
    scene.actionManager.registerAction(
      new BABYLON.ExecuteCodeAction(
        BABYLON.ActionManager.OnKeyUpTrigger,
        evt => {
          EventBus.$emit("keyup", evt.sourceEvent.key);
        }
      )
    );

    scene.onPointerMove = function (evt, result) {
        const pickResult = scene.pick(evt.offsetX, evt.offsetY);
        if (pickResult.hit) {
          EventBus.$emit("scene-pointer-moved", {
            x: pickResult.pickedPoint.x,
            y: pickResult.pickedPoint.z,
          })
        }
    }

    scene.onPointerDown = function castRay(e) {
      var ray = scene.createPickingRay(scene.pointerX, scene.pointerY, BABYLON.Matrix.Identity(), scene.activeCamera);

      var hit = scene.pickWithRay(ray);
      if (hit.pickedMesh) {
        const gameObject = getMeshRoot(hit.pickedMesh)
        if (gameObject) {
          EventBus.$emit("game-object-clicked", {
            game_object: gameObject.metadata.state.payload,
            x: e.pageX,
            y: e.pageY,
          });
        }
      }

      EventBus.$emit("scene-pointer-down")
    };

    window.addEventListener("resize", () => {
      this.resizeCanvas()
    });
    this.resizeCanvas()
  }

  // Makes objects transparent if character is behind something
  castAlphaRay() {
    const direction = this.scene.activeCamera.getDirection(new BABYLON.Vector3.Forward())
    const ray = new BABYLON.Ray(this.scene.activeCamera.position, direction, 300)
    const hits = this.scene.multiPickWithRay(ray)
    let alphaObjects = []
    let alphaObjectsIds = []

    if (hits) {
      for (var i = 0; i < hits.length; i++){
        const meshRoot = getMeshRoot(hits[i].pickedMesh)

        if (meshRoot && !NON_TRANSPARENT_OBJECT_TYPES.includes(meshRoot.metadata.state.payload.Type )) {
          if (!alphaObjectsIds.includes(meshRoot.metadata.id)) {
            alphaObjects.push(meshRoot)
            alphaObjectsIds.push(meshRoot.metadata.id)
          }
        }
      }
    }

    this.previousAlphaMeshes.forEach((mesh) => {
      if (mesh && mesh.metadata && !alphaObjectsIds.includes(mesh.metadata.id)) {
        removeAlpha(mesh, this.scene)
      }
    })

    alphaObjects.forEach((mesh) => {
      if (!this.previousAlphaMeshesIds.includes(mesh.metadata.id)) {
        addAlpha(mesh, this.scene)
      }
    })

    this.previousAlphaMeshes = alphaObjects
    this.previousAlphaMeshesIds = alphaObjectsIds
  }

  resizeCanvas() {
    this.canvas.width = document.body.clientWidth;
    this.canvas.height = document.body.clientHeight;
    this.engine.resize();
  }

  createObjects() {
    this.light = new Light(this.scene);
    this.light.create();

    this.grid = new Grid(this.scene);
    //this.grid.create();

    showWorldAxis(1, this.scene);
  }

  runRenderLoop() {
    this.lastTick = Date.now()
    this.engine.runRenderLoop(() => {
      if (this.scene.activeCamera) {
        this.scene.render();
      }
      const ms = Date.now();
      this.renderObservers.forEach(observer => {
        observer.obj.update(ms - this.lastTick);
      });
      this.lastTick = ms;
      if (this.fpsEl) {
        this.fpsEl.innerHTML = this.engine.getFps().toFixed() + " fps";
      } else {
        this.fpsEl = document.getElementById("fps").firstChild;
      }
    });
  }

  addRenderObserver(id, observer) {
    this.renderObservers.push({
      id: id,
      obj: observer
    });
  }

  removeRenderObserver(id) {
    this.renderObservers = this.renderObservers.filter(ob => ob.id !== id);
  }
}

const gameObserver = new GameObserver();

export default gameObserver;
