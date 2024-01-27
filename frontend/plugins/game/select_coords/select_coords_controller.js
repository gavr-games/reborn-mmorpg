import SelectCoordsObserver from "~/plugins/game/select_coords/select_coords_observer";
import { EventBus } from "~/plugins/game/event_bus";

const IDLE_STATE = 0
const SELECT_COORDS_AND_ROTATION_STATE = 1
const STICK_TO_GRID = 0.5

class SelectCoordsController {
  constructor() {
    this.observer = new SelectCoordsObserver(this.state);
    this.x = 0
    this.y = 0
    this.rotation = 0
    this.itemKey = null
    this.item = null
    this.cmd = null
    this.state = IDLE_STATE
    this.selectionHandler = params => {
      this.state = SELECT_COORDS_AND_ROTATION_STATE
      this.itemKey = params.item_key
      this.item = params.item
      this.callback = params.callback
      this.observer.create(params.item_key, this.x, this.y)
    };
    this.pointerMovedHandler = params => {
      if (this.state === SELECT_COORDS_AND_ROTATION_STATE && params) {
        this.x = params.x - params.x % STICK_TO_GRID
        this.y = params.y - params.y % STICK_TO_GRID
        this.observer.update(this.x, this.y)
      }
    };
    this.pointerDownHandler = () => {
      if (this.state === SELECT_COORDS_AND_ROTATION_STATE) {
        this.state = IDLE_STATE
        this.observer.remove()
        // x and y transformation is required because in engine upper left corner is stored as x/y,
        // but on frontend all assets have pivot points in the center of the object.
        const transformX = this.rotation == 0 ? this.item.width / 2 : this.item.height / 2
        const transformY = this.rotation == 0 ? this.item.height / 2 : this.item.width / 2
        this.callback(this.x - transformX, this.y - transformY, this.rotation)
        this.rotation = 0
      }
    };
    this.keyDownHandler = (key) => {
      if (this.state === SELECT_COORDS_AND_ROTATION_STATE) {
        switch (key) {
          case "Escape":
            this.state = IDLE_STATE
            this.observer.remove()
            this.rotation = 0
            break;
          case "ArrowRight":
          case "ArrowLeft":
            this.rotation = this.rotation == 0 ? Math.PI / 2 : 0
            this.observer.rotate()
            break;
        }
      }
    }
    EventBus.$on("select-coords-and-rotation", this.selectionHandler)
    EventBus.$on("scene-pointer-moved", this.pointerMovedHandler)
    EventBus.$on("scene-pointer-down", this.pointerDownHandler)
    EventBus.$on("keydown", this.keyDownHandler)
    // rotation pressed
  }

  remove() {
    this.observer.remove()
    this.rotation = 0
    EventBus.$off("select-coords-and-rotation", this.selectionHandler)
    EventBus.$off("scene-pointer-moved", this.pointerMovedHandler)
    EventBus.$off("scene-pointer-moved", this.pointerDownHandler)
    EventBus.$off("keydown", this.keyDownHandler)
  }
}

export default SelectCoordsController;
