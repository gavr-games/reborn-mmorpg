import MobObserver from '~/plugins/game/objects/mob/mob_observer'
import MobState from '~/plugins/game/objects/mob/mob_state'
import { EventBus } from '~/plugins/game/event_bus'

class MobController {
  constructor (gameObject) {
    this.state = new MobState(gameObject)
    this.observer = new MobObserver(this.state)
    this.performMeleeHit = (data) => {
      if (data.object.Id === this.state.id) {
        this.observer.meleeHit(data.weapon)
      }
    }
    EventBus.$on('melee_hit_attempt', this.performMeleeHit)
  }

  update (gameObject) {
    this.state.update(gameObject)
    this.observer.updatePosition()
  }

  remove () {
    EventBus.$off('melee_hit_attempt', this.performMeleeHit)
    this.state = null
    this.observer.remove()
  }

  selectAsTarget () {
    this.observer.selectAsTarget()
  }

  deselectAsTarget () {
    this.observer.deselectAsTarget()
  }
}

export default MobController
