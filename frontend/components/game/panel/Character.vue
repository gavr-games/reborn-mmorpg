<template>
  <GameDraggablePanel :panelId="'character'">
    <div v-if="showCharacterInfoPanel" id="character-info-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4>Character</h4>
        <div v-for="(slotItem, slotKey) in characterInfo.slots" :key="slotKey">
          <span>{{ slotKey }}:</span>
          <span v-if="slotItem">
            <GameItem v-bind:item="slotItem" />
          </span>
        </div>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showCharacterInfoPanel: false,
      characterInfo: {
        slots: {}
      }
    }
  },

  created () {
    EventBus.$on('character_info', this.showCharacterInfo)
    EventBus.$on('equip_item', this.equipItem)
    EventBus.$on('unequip_item', this.unequipItem)
  },

  beforeDestroy () {
    EventBus.$off('character_info')
    EventBus.$off('equip_item', this.equipItem)
    EventBus.$off('unequip_item', this.unequipItem)
  },

  methods: {
    showCharacterInfo (data) {
      this.showCharacterInfoPanel = true
      this.characterInfo = data
      localStorage.setItem('open_character_info', 'true')
    },
    equipItem (data) {
      if (this.characterInfo.id === data.character_id) {
        this.characterInfo.slots[data.slot] = data.item
      }
    },
    unequipItem (data) {
      if (this.characterInfo.id === data.character_id) {
        this.characterInfo.slots[data.slot] = null
      }
    },
    close () {
      this.showCharacterInfoPanel = false
      localStorage.removeItem('open_character_info')
    }
  }
}
</script>

<style lang="scss">
#character-info-panel {
}
</style>
