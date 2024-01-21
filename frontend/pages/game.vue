<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="game-panel"><div class="game-panel-content">0</div></div>
    <GameCharacterMenu />
    <GamePanelCraft />
    <GamePanelNpcTrade />
    <GamePanelItemInfo />
    <GamePanelCharacter />
    <GamePanelMap />
    <GamePanelGm />
    <GamePanelTarget />
    <GamePanelEffects />
    <div v-for="(container, _key) in gameContainers" :key="container.id">
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
          EventBus.$emit("keydown", event.key);
        }
      })
      window.addEventListener("keyup",function (event) {
        if (document.activeElement && document.activeElement.tagName != "INPUT") {
          EventBus.$emit("keyup", event.key);
        }
      })
      document.getElementById('game-canvas').addEventListener("mouseover",function (event) {
        document.getElementById("game-canvas").focus()
      })
    }
  },

  created() {
    EventBus.$on("container_items", this.addContainer)
    EventBus.$on("init_game", this.initGame)
  },

  beforeDestroy() {
    this.$game.destroy()
    EventBus.$off("container_items", this.addContainer)
    EventBus.$off("init_game", this.initGame)
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
    },
    initGame() {
      // Init containers
      let openedContainers = localStorage.getItem("opened_containers")
      if (openedContainers) {
        openedContainers = JSON.parse(openedContainers)
        openedContainers.forEach(containerId => {
          EventBus.$emit("perform-game-action", {
            cmd: "open_container",
            params: containerId
          })
        })
      }
      // Init Character info
      const openCharacterInfo = localStorage.getItem("open_character_info")
      if (openCharacterInfo) {
        EventBus.$emit("perform-game-action", {
          cmd: "get_character_info",
          params: {}
        });
      }
    }
  }
}
</script>
