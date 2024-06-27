<template>
  <div
    draggable
    class="draggable-panel"
    :style="{ top: top + 'px', left: left + 'px' }"
    @dragstart="startDrag($event)"
    @dragend="endDrag($event)"
  >
    <slot />
  </div>
</template>

<script>
export default {
  props: ['panelId'],
  data () {
    return {
      top: 0,
      left: 0,
      startPageX: 0,
      startPageY: 0
    }
  },

  mounted () {
    const pos = localStorage.getItem(`panel-${this.panelId}`)
    if (pos) {
      const coords = JSON.parse(pos)
      this.left = coords.left
      this.top = coords.top
    }
  },

  methods: {
    startDrag (evt) {
      this.startPageX = evt.pageX
      this.startPageY = evt.pageY
      evt.dataTransfer.dropEffect = 'move'
      evt.dataTransfer.effectAllowed = 'move'
    },
    endDrag (evt) {
      this.left = this.left + (evt.pageX - this.startPageX)
      this.top = this.top + (evt.pageY - this.startPageY)
      localStorage.setItem(`panel-${this.panelId}`, JSON.stringify({ left: this.left, top: this.top }))
    }
  }
}
</script>

<style lang="scss">
.draggable-panel {
  position: absolute;
}
</style>
