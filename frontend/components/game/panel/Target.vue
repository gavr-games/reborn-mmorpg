<template>
  <div id="target-info-panel" class="rpgui-container framed-golden" v-if="showTargetInfoPanel">
    <h4>Target</h4>
    <div v-if="targetInfo">
      {{ targetInfo["kind"] }} <br />
      <a href="#" @click="triggerDeselect">Deselect</a>
    </div>
    <div id="fighting-panel">
      <span>1</span> to Hit
    </div>
  </div>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showTargetInfoPanel: false,
      targetInfo: null
    }
  },

  created () {
    EventBus.$on('select_target', this.selectTarget)
    EventBus.$on('deselect_target', this.deselectTarget)
  },

  beforeDestroy () {
    EventBus.$off('select_target', this.selectTarget)
    EventBus.$off('deselect_target', this.deselectTarget)
  },

  methods: {
    selectTarget (data) {
      this.showTargetInfoPanel = true
      this.targetInfo = data
    },
    deselectTarget () {
      this.showTargetInfoPanel = false
      this.targetInfo = null
    },
    triggerDeselect () {
      EventBus.$emit('perform-game-action', {
        cmd: 'deselect_target',
        params: this.targetInfo.Id
      })
    }
  }
}
</script>

<style lang="scss">
#target-info-panel {
  position: absolute;
  top: 0px;
  right: 100px;
  width: 200px;
}
</style>
