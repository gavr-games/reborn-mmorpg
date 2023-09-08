<template>
  <div id="character-info-panel" class="rpgui-container framed-golden" v-if="showCharacterInfoPanel">
    <h4>Character</h4>
    <div v-for="(slotItem, slotKey) in characterInfo.slots" :key="slotKey">
      <span>{{ slotKey }}:</span>
      <span v-if="slotItem">
        <GameItem v-bind:item="slotItem" />
      </span>
    </div>
    <button type="button" class="rpgui-button" @click="showCharacterInfoPanel = false"><p>Close</p></button>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showCharacterInfoPanel: false,
      characterInfo: {
        slots: {},
      },
    }
  },

  created() {
    EventBus.$on("character_info", this.showCharacterInfo)
  },

  beforeDestroy() {
    EventBus.$off("character_info")
  },

  methods: {
    showCharacterInfo(data) {
      this.showCharacterInfoPanel = true
      this.characterInfo = data
    },
  }
}
</script>


<style>
#character-info-panel {
  position: absolute;
  top: 50px;
  left: 50px;
}
</style>
