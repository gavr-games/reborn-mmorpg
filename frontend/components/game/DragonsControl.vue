<template>
  <div id="dragons-control-menu">
    <div v-show="showDragonsList" id="dragons-control-list" class="game-panel">
      <div class="game-panel-content">
        <div v-for="dragon in dragons" :key="dragon.id">
          <input v-model="selectedDragons" :value="dragon.id" class="rpgui-checkbox" type="checkbox" data-rpguitype="checkbox">
          <label @click="toggleDragonSelected(dragon.id)">{{ dragon.kind }} Lvl. {{ dragon.level }}</label>
        </div>
      </div>
    </div>
    <div class="game-panel">
      <div class="game-panel-content">
        <div class="menu-item" @click="toggleDragonsList">
          {{ showDragonsList ? "<" : ">" }}
        </div>
        <div class="menu-item attack-icon" title="Attack My Target" @click="attackMyTarget" />
        <div class="menu-item follow-icon" title="Follow Me" @click="followMe" />
        <div class="menu-item stop-icon" title="Stop" @click="stop" />
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showDragonsList: false,
      dragons: [],
      selectedDragons: []
    }
  },

  created () {
    EventBus.$on('dragons_info', this.setDragonsList)
  },

  mounted () {
    const selectedControlDragons = localStorage.getItem('selected_control_dragons')
    if (selectedControlDragons) {
      this.selectedDragons = JSON.parse(selectedControlDragons)
    }
  },

  beforeDestroy () {
    EventBus.$off('dragons_info', this.setDragonsList)
  },

  methods: {
    attackMyTarget () {
      this.selectedDragons.forEach((dragonId) => {
        EventBus.$emit('perform-game-action', {
          cmd: 'attack_my_target',
          params: dragonId
        })
      })
    },
    followMe () {
      this.selectedDragons.forEach((dragonId) => {
        EventBus.$emit('perform-game-action', {
          cmd: 'follow',
          params: dragonId
        })
      })
    },
    stop () {
      this.selectedDragons.forEach((dragonId) => {
        EventBus.$emit('perform-game-action', {
          cmd: 'order_to_stop',
          params: dragonId
        })
      })
    },
    setDragonsList (data) {
      this.dragons = data.dragons
      let selectedControlDragons = localStorage.getItem('selected_control_dragons')
      if (selectedControlDragons) {
        this.selectedDragons = []
        selectedControlDragons = JSON.parse(selectedControlDragons)
        selectedControlDragons.forEach((dragonId) => {
          if (this.dragons.some(dragon => dragon.id === dragonId)) {
            this.selectedDragons.push(dragonId)
          }
        })
        localStorage.setItem('selected_control_dragons', JSON.stringify(this.selectedDragons))
      }
    },
    toggleDragonsList () {
      this.showDragonsList = !this.showDragonsList
      if (this.showDragonsList) {
        EventBus.$emit('perform-game-action', {
          cmd: 'get_dragons_info',
          params: {}
        })
      }
    },
    toggleDragonSelected (dragonId) {
      const addOrRemove = (arr, item) => arr.includes(item) ? arr.filter(i => i !== item) : [...arr, item]
      this.selectedDragons = addOrRemove(this.selectedDragons, dragonId)
      localStorage.setItem('selected_control_dragons', JSON.stringify(this.selectedDragons))
    }
  }
}
</script>

<style lang="scss">
#dragons-control-menu {
  position: absolute;
  top: 70px;
  left: 0;
  .menu-item {
    display: block;
    width: 32px;
    height: 32px;
    padding-bottom: 5px;
    opacity: 0.8;
    text-align: center;
    &:hover {
      opacity: 1.0;
      cursor: url("~assets/img/cursor/point.png") 10 0, auto;
    }
  }
  .attack-icon {
    background-image: url("~assets/img/icons/dragons-attack.png");
    background-repeat: no-repeat;
  }
  .follow-icon {
    background-image: url("~assets/img/icons/dragons-follow.png");
    background-repeat: no-repeat;
  }
  .stop-icon {
    background-image: url("~assets/img/icons/dragons-stop.png");
    background-repeat: no-repeat;
  }
}
#dragons-control-list {
  position: absolute;
  left: 80px;
}
</style>
