<template>
  <a href="#" class="close-btn" @click="closeCallback()" :style="`top: ${yPos}px; right: ${xPos}px;`" />
</template>

<script>
export default {
  props: ['closeCallback', 'x', 'y'],
  data () {
    return {
      xPos: -8,
      yPos: -8
    }
  },
  mounted () {
    if (this.x) {
      this.xPos = this.x
    }
    if (this.y) {
      this.yPos = this.y
    }
  }
}
</script>

<style lang="scss">
.close-btn {
  position: absolute;
  top: -8px;
  right: -8px;
  width: 16px;
  height: 16px;
  background-image: url("~assets/img/close.png");
  &:hover {
    transition: rotate 1s;
  }
}
</style>
