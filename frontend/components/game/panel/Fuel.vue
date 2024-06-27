<template>
  <GameDraggablePanel :panelId="'fuel'">
    <div v-if="showFuelPanel" id="fuel-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4 class="heading">
          Add fuel
        </h4>
        <div v-if="item['Properties']">
          Fuel: {{ item["Properties"]["fuel"] }} / {{ item["Properties"]["max_fuel"] }}
        </div>
        <div>Drop fuel here:</div>
        <div class="empty-slot" @dragover="allowDrag" @drop="onDrop($event)" />
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showFuelPanel: false,
      item: {},
      callback: null
    }
  },

  created () {
    EventBus.$on('select-fuel-id', this.showFeed)
  },

  beforeDestroy () {
    EventBus.$off('select-fuel-id', this.showFeed)
  },

  methods: {
    showFeed (data) {
      this.showFuelPanel = true
      this.item = data.item
      this.callback = data.callback
    },
    onDrop (evt) {
      evt.stopPropagation()
      const itemId = evt.dataTransfer.getData('item_id')
      this.callback(itemId)
    },
    allowDrag (evt) {
      evt.preventDefault()
    },
    close () {
      this.showFuelPanel = false
    }
  }
}
</script>

<style lang="scss">
#fuel-panel {
  color: white;
  .heading {
    margin-top: 0px;
  }
  .empty-slot {
    width: 32px;
    height: 32px;
    border: 1px solid white;
    margin-left: 32px;
  }
}
.game-panel-content {
  div {
    color: white;
    font-size: 11px;
  }
}
</style>
