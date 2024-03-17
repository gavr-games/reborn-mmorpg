<template>
  <div class="rpgui-progress current-action-bar" data-rpguitype="progress" v-if="showActionBar">
    <div class="rpgui-progress-track">
      <div class="rpgui-progress-fill" :style="`left: 0px; width: ${progress}%;`"></div>
    </div>
    <div class=" rpgui-progress-left-edge"></div>
    <div class=" rpgui-progress-right-edge"></div>
  </div>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {

  data () {
    return {
      progress: 0,
      showActionBar: false,
      interval: null,
      currentPlayerId: 0
    }
  },

  created () {
    this.currentPlayerId = this.$store.state.characters.selectedCharacterId
    EventBus.$on('start_delayed_action', this.startAction)
    EventBus.$on('cancel_delayed_action', this.cancelAction)
    EventBus.$on('finish_delayed_action', this.finishAction)
  },

  beforeDestroy () {
    EventBus.$off('start_delayed_action', this.startAction)
    EventBus.$off('cancel_delayed_action', this.cancelAction)
    EventBus.$off('finish_delayed_action', this.finishAction)
  },

  methods: {
    startAction (data) {
      if (data.object.player_id && data.object.player_id === this.currentPlayerId) {
        this.showActionBar = true
        this.progress = 0
        this.interval = setInterval(() => {
          if (this.progress < 100) {
            this.progress += 2
          }
        }, data.duration / 50)
      }
    },
    cancelAction(data) {
      if (data.object.player_id && data.object.player_id === this.currentPlayerId) {
        this.showActionBar = false
        if (this.interval != null) {
          clearInterval(this.interval)
          this.interval = null
        }
      }
    },
    finishAction(data) {
      if (data.object.player_id && data.object.player_id === this.currentPlayerId) {
        this.showActionBar = false
        if (this.interval != null) {
          clearInterval(this.interval)
          this.interval = null
        }
      }
    }
  }
}
</script>


<style>
.current-action-bar {
  position: absolute;
  width: 250px;
  z-index: 999;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  margin-top: 10px;
}
</style>
