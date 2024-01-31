import SurfaceObserver from '~/plugins/game/objects/surface/surface_observer'
import SurfaceState from '~/plugins/game/objects/surface/surface_state'

class SurfaceController {
  constructor (gameObject) {
    this.state = new SurfaceState(gameObject)
    this.observer = new SurfaceObserver(this.state)
  }

  remove () {
    this.state = null
    this.observer.remove()
  }
}

export default SurfaceController
