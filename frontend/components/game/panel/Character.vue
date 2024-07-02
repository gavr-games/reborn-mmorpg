<template>
  <GameDraggablePanel :panelId="'character'">
    <div v-if="showCharacterInfoPanel" id="character-info-panel" class="game-panel">
      <div id="character">
        <GameCloseIcon :close-callback="close" />
        <div class="game-panel-content">
          <div v-for="(slotItem, slotKey) in characterInfo.slots" :id="slotKey" :key="slotKey" class="slot" :title="slotKey" @dragover="allowDrag" @drop="onDrop">
            <span v-if="slotItem">
              <GameItem v-bind:item="slotItem" />
            </span>
          </div>
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
    onDrop (evt) {
      evt.stopPropagation()
      const itemID = evt.dataTransfer.getData('item_id')
      EventBus.$emit('perform-game-action', {
        cmd: 'equip_item',
        params: itemID
      })
    },
    allowDrag (evt) {
      evt.preventDefault()
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
  #character {
    width: 140px;
    height: 314px;
    background-image: url("~assets/img/character-bg.png");
    background-size: contain;
    background-repeat: no-repeat;
    .slot {
      width: 32px;
      height: 32px;
      border: 1px solid white;
      position: absolute;
    }
    #back {
      left: 126px;
      top: 48px;
    }
    #body {
      left: 76px;
      top: 108px;
    }
    #left_arm {
      left: 124px;
      top: 198px;
    }
    #right_arm {
      left: 24px;
      top: 198px;
    }
  }
}
</style>
