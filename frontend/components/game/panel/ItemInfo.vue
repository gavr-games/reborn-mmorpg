<template>
  <GameDraggablePanel :panelId="'item_info'">
    <div v-if="showItemInfoPanel" id="item_info-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4 class="heading">
          {{ itemInfo["kind"] }}
        </h4>
        <div v-if="itemInfo['crafted_by']">
          Crafted by: {{ itemInfo["crafted_by"]["name"] }}
        </div>
        <div v-if="itemInfo['payed_until']">
          Payed until: {{ new Date(itemInfo["payed_until"]) }}
        </div>
        <div v-if="itemInfo['level']">
          Level: {{ itemInfo["level"] }}
        </div>
        <div v-if="itemInfo['experience']">
          Experience: {{ itemInfo["experience"] }}
        </div>
        <div v-if="itemInfo['fullness']">
          Fullness: {{ itemInfo["fullness"] }}/{{ itemInfo["max_fullness"] }}
        </div>
        <div v-if="itemInfo['fuel']">
          Fuel: {{ itemInfo["fuel"] }}/{{ itemInfo["max_fuel"] }}
        </div>
        <div v-if="itemInfo['health']">
          Health: {{ itemInfo["health"] }}/{{ itemInfo["max_health"] }}
        </div>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showItemInfoPanel: false,
      itemInfo: {}
    }
  },

  created () {
    EventBus.$on('item_info', this.showItemInfo)
  },

  beforeDestroy () {
    EventBus.$off('item_info', this.showItemInfo)
  },

  methods: {
    showItemInfo (data) {
      this.showItemInfoPanel = true
      this.itemInfo = data
    },
    close () {
      this.showItemInfoPanel = false
    }
  }
}
</script>

<style lang="scss">
#item_info-panel {
  color: white;
  .heading {
    margin-top: 0px;
  }
}
.game-panel-content {
  div {
    color: white;
    font-size: 11px;
  }
}
</style>
