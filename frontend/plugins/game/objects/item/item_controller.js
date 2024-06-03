import ItemObserver from '~/plugins/game/objects/item/item_observer'
import ItemState from '~/plugins/game/objects/item/item_state'

class ItemController {
  constructor (gameObject) {
    this.state = new ItemState(gameObject)
    this.observer = new ItemObserver(this.state)
  }

  update (gameObject) {
    const oldItemState = this.state.state
    this.state.update(gameObject)
    this.observer.updateLifted()
    // Change item model (open/close door, burn/shut fireplace)
    if (oldItemState !== this.state.state) {
      this.observer.changeModel()
    }
  }

  remove () {
    this.state = null
    this.observer.remove()
  }
}

export default ItemController
