import * as BABYLON from 'babylonjs'
import * as Loaders from 'babylonjs-loaders'
import Atlas from '~/plugins/game/atlas/atlas'
import Loader from '~/plugins/game/atlas/loader'
import Light from '~/plugins/game/light/light'
import { EventBus } from '~/plugins/game/event_bus'
import showWorldAxis from '~/plugins/game/utils/world_axis'
import Grid from '~/plugins/game/utils/grid'
import getMeshRoot from '~/plugins/game/utils/get_mesh_root'
import addAlpha from '~/plugins/game/utils/add_alpha'
import removeAlpha from '~/plugins/game/utils/remove_alpha'

const NON_TRANSPARENT_OBJECT_TYPES = ['player', 'surface']
const INTERRACTABLE_OBJECT_ALPHA = 0.8

const LEFT_MOUSE_BUTTON = 0
const RIGHT_MOUSE_BUTTON = 2

class GameObserver {
  constructor () {
    this.canvas = null
    this.engine = null
    this.scene = null
    this.loader = null
    this.light = null
    this.fpsEl = null
    this.grid = null
    this.lastTick = 0
    this.loaded = false
    this.renderObservers = []
    this.previousAlphaMeshes = []
    this.previousAlphaMeshesIds = []
    this.pointerOverMesh = null // highlight interractable objects
    this.pointerOverMeshId = null
    this.pickCoordsPlane = null // this plane is used for picking coords on pointer moved
  }

  init () {
    this.canvas = document.getElementById('game-canvas')
    this.engine = new BABYLON.Engine(this.canvas, true, {
      preserveDrawingBuffer: true,
      stencil: true
    })
    Loaders.OBJFileLoader.OPTIMIZE_WITH_UV = true
    this.createScene()
  }

  createScene () {
    this.scene = new BABYLON.Scene(this.engine, {
      useGeometryUniqueIdsMap: true,
      useMaterialMeshMap: true,
      useClonedMeshMap: true
    })
    this.scene.actionManager = new BABYLON.ActionManager(this.scene)
    this.registerActions(this.scene)

    this.scene.useGeometryIdsMap = true
    this.scene.useMaterialMeshMap = true
    this.scene.useClonedMeshMap = true

    // this.scene.performancePriority = BABYLON.ScenePerformancePriority.Aggressive
    this.scene.clearColor = BABYLON.Color3.FromInts(66, 245, 102) // default grass color for world
    this.createObjects()

    this.loader = new Loader(
      this.scene,
      () => {
        this.runRenderLoop()
        this.scene.registerBeforeRender(() => {
          this.castAlphaRay()
        })
        // this.scene.debugLayer.show()
        this.loaded = true
        EventBus.$emit('scene-created', this.scene, this.canvas)
      },
      Atlas
    )
    this.loader.load()
  }

  registerActions (scene) {
    const gameObserver = this
    scene.onPointerMove = function (evt, result) {
      const ray = scene.createPickingRay(scene.pointerX, scene.pointerY, BABYLON.Matrix.Identity(), null)
      const hit = scene.pickWithRay(ray)
      const pickedPoint = hit.pickedPoint
      if (pickedPoint) {
        EventBus.$emit('scene-pointer-moved', {
          x: pickedPoint.x,
          y: pickedPoint.z
        })
      }
      if (hit.pickedMesh) {
        const gameObject = getMeshRoot(hit.pickedMesh)
        if (gameObject && gameObject.metadata && gameObject.metadata.state.payload.Properties.actions !== undefined) {
          if (gameObserver.pointerOverMeshId !== gameObject.metadata.state.payload.Properties.id) {
            gameObserver.removePointerOverMesh()
            addAlpha(gameObject, gameObserver.scene, INTERRACTABLE_OBJECT_ALPHA)
            gameObserver.pointerOverMesh = gameObject
            gameObserver.pointerOverMeshId = gameObject.metadata.state.payload.Properties.id
          }
        } else {
          gameObserver.removePointerOverMesh()
        }
      } else {
        gameObserver.removePointerOverMesh()
      }
    }

    scene.onPointerDown = function castRay (e) {
      const ray = scene.createPickingRay(scene.pointerX, scene.pointerY, BABYLON.Matrix.Identity(), scene.activeCamera)

      const hit = scene.pickWithRay(ray)
      if (hit.pickedMesh) {
        if (hit.pickedMesh.id === 'xy-coords-plane') {
          const pickedPoint = hit.pickedPoint
          if (scene.getNodeByName('select-coords-item') === null && e.button === LEFT_MOUSE_BUTTON) {
            EventBus.$emit('perform-game-action', {
              cmd: 'move_xy',
              params: {
                x: pickedPoint.x,
                y: pickedPoint.z
              }
            })
          }
        } else {
          const gameObject = getMeshRoot(hit.pickedMesh)
          if (gameObject) {
            if (e.button === LEFT_MOUSE_BUTTON) {
              EventBus.$emit('game-object-clicked', {
                game_object: gameObject.metadata.state.payload,
                x: e.pageX,
                y: e.pageY
              })
            } else if (e.button === RIGHT_MOUSE_BUTTON) {
              EventBus.$emit('game-object-right-clicked', {
                game_object: gameObject.metadata.state.payload,
                x: e.pageX,
                y: e.pageY
              })
            }
          }
        }
      }

      EventBus.$emit('scene-pointer-down')
    }

    window.addEventListener('resize', () => {
      this.resizeCanvas()
    })
    this.resizeCanvas()
  }

  removePointerOverMesh () {
    if (this.pointerOverMeshId !== null) {
      removeAlpha(this.pointerOverMesh, this.scene, INTERRACTABLE_OBJECT_ALPHA)
      this.pointerOverMeshId = null
      this.pointerOverMesh = null
    }
  }

  // Makes objects transparent if character is behind something
  castAlphaRay () {
    const direction = this.scene.activeCamera.getDirection(new BABYLON.Vector3.Forward())
    const ray = new BABYLON.Ray(this.scene.activeCamera.position, direction, 300)
    const hits = this.scene.multiPickWithRay(ray)
    const alphaObjects = []
    const alphaObjectsIds = []

    if (hits) {
      for (let i = 0; i < hits.length; i++) {
        const meshRoot = getMeshRoot(hits[i].pickedMesh)

        if (meshRoot && !NON_TRANSPARENT_OBJECT_TYPES.includes(meshRoot.metadata.state.payload.Type)) {
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

  resizeCanvas () {
    this.canvas.width = document.body.clientWidth
    this.canvas.height = document.body.clientHeight
    this.engine.resize()
  }

  createObjects () {
    this.light = new Light(this.scene)
    this.light.create()

    this.grid = new Grid(this.scene)
    // this.grid.create();

    // Create plane to pick coords - replace 1000 with area width
    this.pickCoordsPlane = BABYLON.MeshBuilder.CreatePlane('xy-coords-plane', { height: 1000, width: 1000 }, this.scene)
    this.pickCoordsPlane.position.x = 0
    this.pickCoordsPlane.position.z = 0
    this.pickCoordsPlane.position.y = 0.05
    this.pickCoordsPlane.rotate(BABYLON.Axis.X, Math.PI / 2)
    this.pickCoordsPlane.material = new BABYLON.StandardMaterial('ObjectplaneMaterial', this.scene)
    this.pickCoordsPlane.material.alpha = 0
    this.pickCoordsPlane.convertToUnIndexedMesh()
    this.pickCoordsPlane.doNotSyncBoundingInfo = true

    showWorldAxis(1, this.scene)
  }

  runRenderLoop () {
    this.lastTick = Date.now()
    this.engine.runRenderLoop(() => {
      if (this.scene.activeCamera) {
        this.scene.render()
      }
      const ms = Date.now()
      const tickDelta = ms - this.lastTick
      this.renderObservers.forEach((observer) => {
        observer.obj.update(tickDelta)
      })
      this.lastTick = ms
      if (this.fpsEl) {
        this.fpsEl.innerHTML = this.engine.getFps().toFixed() + ' fps'
      } else {
        this.fpsEl = document.getElementById('fps').firstChild
      }
    })
  }

  addRenderObserver (id, observer) {
    this.renderObservers.push({
      id,
      obj: observer
    })
  }

  removeRenderObserver (id) {
    this.renderObservers = this.renderObservers.filter(ob => ob.id !== id)
  }
}

const gameObserver = new GameObserver()

export default gameObserver
