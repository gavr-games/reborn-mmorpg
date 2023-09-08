<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="rpgui-container framed-golden">0</div>
    <GameCharacterMenu />
    <GamePanelCharacter />
    <GameChat />
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
    }
  },

  mounted() {
    if (!this.$auth.loggedIn || !this.$store.state.characters.selectedCharacterId) {
      this.$router.push('login')
    } else {
      this.$game.init(this.$auth.strategy.token.get(), this.$store.state.characters.selectedCharacterId)
    }
  },

  beforeDestroy() {
    this.$game.destroy()
  },

  computed: {
    charId() {
      return this.$store.state.characters.selectedCharacterId
    }
  },

  methods: {
  }
}
</script>
