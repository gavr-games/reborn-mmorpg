<template>
  <div class="rpgui-content" style="overflow-y:scroll;" @dragover="allowDrag" @drop="onDrop($event, key)">
    <canvas id="game-canvas" />
    <div id="fps" class="game-panel">
      <div class="game-panel-content">
        0
      </div>
    </div>
    <GameCharacterMenu />
    <GameExperienceBar />
    <GameDragonsControl />
    <GamePanelCraft />
    <GamePanelDragons />
    <GamePanelFeed />
    <GamePanelFuel />
    <GamePanelDungeons />
    <GamePanelNpcTrade />
    <GamePanelItemInfo />
    <GamePanelCharacter />
    <GamePanelMap />
    <GamePanelGm />
    <GamePanelTarget />
    <GamePanelEffects />
    <div v-for="(container) in gameContainers" :key="container.id">
      <GamePanelContainer :container="container" />
    </div>
    <GameObjectContextMenu />
    <GameCurrentActionBar />
    <GameChat />
  </div>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      gameContainers: []
    }
  },

  computed: {
    charId () {
      return this.$store.state.characters.selectedCharacterId
    }
  },

  mounted () {
    if (!this.$auth.loggedIn || !this.$store.state.characters.selectedCharacterId) {
      this.$router.push('login')
    } else {
      this.$game.init(this.$auth.strategy.token.get(), this.$store.state.characters.selectedCharacterId)
      // Give babylon scene back the keyboard control
      window.addEventListener('keydown', function (event) {
        if (document.activeElement && document.activeElement.tagName !== 'INPUT') {
          EventBus.$emit('keydown', event.key)
        }
      })
      window.addEventListener('keyup', function (event) {
        if (document.activeElement && document.activeElement.tagName !== 'INPUT') {
          EventBus.$emit('keyup', event.key)
        }
      })
      document.getElementById('game-canvas').addEventListener('mouseover', function (event) {
        document.getElementById('game-canvas').focus()
      })
    }
  },

  created () {
    EventBus.$on('container_items', this.addContainer)
    EventBus.$on('init_game', this.initGame)
  },

  beforeDestroy () {
    this.$game.destroy()
    EventBus.$off('container_items', this.addContainer)
    EventBus.$off('init_game', this.initGame)
  },

  methods: {
    addContainer (data) {
      const contIndex = this.gameContainers.findIndex(cont => cont.id === data.id)
      if (contIndex !== -1) {
        this.gameContainers.pop[contIndex] // eslint-disable-line no-unused-expressions
      }
      this.gameContainers.push(data)
    },
    initGame () {
      // Init containers
      let openedContainers = localStorage.getItem('opened_containers')
      if (openedContainers) {
        openedContainers = JSON.parse(openedContainers)
        openedContainers.forEach((containerId) => {
          EventBus.$emit('perform-game-action', {
            cmd: 'open_container',
            params: containerId
          })
        })
      }
      // Init Character info
      const openCharacterInfo = localStorage.getItem('open_character_info')
      if (openCharacterInfo) {
        EventBus.$emit('perform-game-action', {
          cmd: 'get_character_info',
          params: {}
        })
      }
    },
    onDrop (evt, pos) {
      evt.stopPropagation()
    },
    allowDrag (evt) {
      evt.preventDefault()
    }
  }
}
</script>
