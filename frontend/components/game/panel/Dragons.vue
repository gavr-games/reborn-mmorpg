<template>
  <GameDraggablePanel :panel-id="'dragons'">
    <div v-if="showDragonsPanel" id="dragons-panel" class="game-panel">
      <div class="game-panel-content">
        <h4 class="heading">
          Dragons (Max {{ maxDragons }})
        </h4>
        <p class="hint">
          To get more dragons build hatchery on your claim.
        </p>
        <div v-for="dragon in dragons" :key="dragon.id">
          <div class="dragon">
            <div class="dragon-title" @click="toggleExpandDragon(dragon.id)">
              {{ dragon.kind }} Lvl. {{ dragon.level }}
            </div>
            <div v-if="expandDragons[dragon.id]">
              <p class="dragon-description">
                {{ dragon.alive ? 'Alive' : 'Dead' }} | Exp. {{ dragon.experience }} / {{ expNextLevel(dragon.level) }}
              </p>
              <button v-if="dragon.alive" type="button" class="rpgui-button" @click="teleportDragon(dragon)">
                <p>Teleport to me</p>
              </button>
              <button v-if="!dragon.alive" type="button" class="rpgui-button" @click="resurrectDragon(dragon)">
                <p>Resurrect ({{ 25 * (dragon.level + 1) }})</p>
              </button>
              <button type="button" class="rpgui-button" @click="releaseDragon(dragon)">
                <p>Release</p>
              </button>
            </div>
          </div>
        </div>
        <button type="button" class="rpgui-button" @click="showDragonsPanel = false">
          <p>Close</p>
        </button>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

const EXP_AMOUNT_DIVIDER = 0.05
const EXP_INCREASE_POWER = 2.0

export default {
  data () {
    return {
      showDragonsPanel: false,
      dragons: [],
      maxDragons: 0,
      expandDragons: {}
    }
  },

  created () {
    EventBus.$on('dragons_info', this.showDragonsInfo)
  },

  beforeDestroy () {
    EventBus.$off('dragons_info', this.showDragonsInfo)
  },

  methods: {
    showDragonsInfo (data) {
      this.showDragonsPanel = true
      this.maxDragons = data.max_dragons
      this.dragons = data.dragons
    },
    toggleExpandDragon (dragonId) {
      if (this.expandDragons[dragonId]) {
        this.expandDragons[dragonId] = false
      } else {
        this.expandDragons[dragonId] = true
      }
      this.$forceUpdate()
    },
    releaseDragon (dragon) {
      EventBus.$emit('perform-game-action', {
        cmd: 'release_dragon',
        params: dragon.id
      })
    },
    resurrectDragon (dragon) {
      EventBus.$emit('perform-game-action', {
        cmd: 'resurrect_dragon',
        params: dragon.id
      })
    },
    teleportDragon (dragon) {
      EventBus.$emit('perform-game-action', {
        cmd: 'teleport_dragon_to_owner',
        params: dragon.id
      })
    },
    expNextLevel (level) {
      return Math.pow((level + 1) / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER) - Math.pow(level / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER)
    }
  }
}
</script>

<style lang="scss">
#dragons-panel {
  color: white;
  .heading {
    margin-top: 0px;
  }
  .hint {
    font-size: 8px;
  }
  .dragon {
    .dragon-title {
      color: white;
      &:hover {
        text-decoration: underline;
      }
    }
    border: 1px solid white;
    padding: 5px;
    .dragon-description {
      color: grey;
      margin: 0px;
      font-size: 10px;
    }
    .actions {
      display: flex;
    }
  }
}
</style>
