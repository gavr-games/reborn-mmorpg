<template>
  <GameDraggablePanel :panelId="'feed'">
    <div v-if="showFeedPanel" id="feed-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4 class="heading">
          Feed
        </h4>
        <div v-if="item['Properties']">
          Fullness: {{ item["Properties"]["fullness"] }} / {{ item["Properties"]["max_fullness"] }}
        </div>
        <div>Drop food here:</div>
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
      showFeedPanel: false,
      item: {},
      callback: null
    }
  },

  created () {
    EventBus.$on('select-food-id', this.showFeed)
  },

  beforeDestroy () {
    EventBus.$off('select-food-id', this.showFeed)
  },

  methods: {
    showFeed (data) {
      this.showFeedPanel = true
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
      this.showFeedPanel = false
    }
  }
}
</script>

<style lang="scss">
#feed-panel {
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
