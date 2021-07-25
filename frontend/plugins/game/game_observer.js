import * as BABYLON from "babylonjs";
import * as Loaders from "babylonjs-loaders";
import Atlas from "~/plugins/game/atlas/atlas";
import Loader from "~/plugins/game/atlas/loader";
import Camera from "~/plugins/game/camera/camera";
import Light from "~/plugins/game/light/light";
import Character from "~/plugins/game/character/character";
import { EventBus } from "~/plugins/game/event_bus";
import showWorldAxis from "~/plugins/game/utils/world_axis";

class GameObserver {
  constructor() {
    this.canvas = null;
    this.engine = null;
    this.scene = null;
    this.loader = null;
    this.camera = null;
    this.light = null;
    this.fpsEl = null;
    this.renderObservers = [];
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
    this.scene = new BABYLON.Scene(this.engine);
    this.scene.actionManager = new BABYLON.ActionManager(this.scene);
    this.registerActions(this.scene);

    this.loader = new Loader(
      this.scene,
      () => {
        this.createObjects();
        this.runRenderLoop();
        EventBus.$emit("scene-created", this.scene);
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

    window.addEventListener("resize", () => {
      this.resizeCanvas()
    });
    this.resizeCanvas()
  }

  resizeCanvas() {
    this.canvas.width = document.body.clientWidth;
    this.canvas.height = document.body.clientHeight;
    this.engine.resize();
  }

  createObjects() {
    let character = new Character(this.scene, this.canvas);
    character.create();

    this.light = new Light(this.scene);
    this.light.create();

    this.camera = new Camera(this.scene, this.canvas, character);
    this.camera.create();

    showWorldAxis(1, this.scene)
  }

  runRenderLoop() {
    this.engine.runRenderLoop(() => {
      if (this.scene.activeCamera) {
        this.scene.render();
      }
      this.renderObservers.forEach(observer => {
        observer.obj.update();
      });
      if (this.fpsEl) {
        this.fpsEl.innerHTML = this.engine.getFps().toFixed() + " fps";
      } else {
        this.fpsEl = document.getElementById("fps");
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
