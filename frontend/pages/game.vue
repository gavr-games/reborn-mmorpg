<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="game-panel"><div class="game-panel-content">0</div></div>
    <GameCharacterMenu />
    <GamePanelCraft />
    <GamePanelNpcTrade />
    <GamePanelCharacter />
    <GamePanelMap />
    <GamePanelTarget />
    <GamePanelEffects />
    <div v-for="(container, key) in gameContainers" :key="key">
      {{key}}
      <GamePanelContainer v-bind:container="container" />
    </div>
    <GameObjectContextMenu />
    <GameCurrentActionBar />
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
      // Give babylon scene back the keyboard control
      window.addEventListener("keydown",function (event) {
        if (document.activeElement && document.activeElement.tagName != "INPUT") {
          document.getElementById("game-canvas").focus()
        }
      })
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
