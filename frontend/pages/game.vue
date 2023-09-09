<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="rpgui-container framed-golden">0</div>
    <GameCharacterMenu />
    <GamePanelCharacter />
    <div v-for="(container, key) in gameContainers" :key="key">
      {{key}}
      <GamePanelContainer v-bind:container="container" />
    </div>

    <GameChat />
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      gameContainers: [],
    }
  },

  mounted() {
    if (!this.$auth.loggedIn || !this.$store.state.characters.selectedCharacterId) {
      this.$router.push('login')
    } else {
      this.$game.init(this.$auth.strategy.token.get(), this.$store.state.characters.selectedCharacterId)
    }
  },

  created() {
    EventBus.$on("container_items", this.addContainer)
  },

  beforeDestroy() {
    this.$game.destroy()
    EventBus.$off("container_items", this.addContainer)
  },

  computed: {
    charId() {
      return this.$store.state.characters.selectedCharacterId
    }
  },

  methods: {
    addContainer(data) {
      const contIndex = this.gameContainers.findIndex((cont) => cont.id === data.id)
      if (contIndex !== -1) {
        this.gameContainers.pop[contIndex]
      }
      this.gameContainers.push(data)
    }
  }
}
</script>
