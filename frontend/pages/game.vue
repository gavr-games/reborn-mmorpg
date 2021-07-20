<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps">0</div>
  </div>
</template>

<script>
export default {
  mounted() {
    this.$game.init()
  },

  computed: {
    charId () {
      return this.$store.state.characters.selectedCharacterId
    }
  },
}
</script>
