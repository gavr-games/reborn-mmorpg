<template>
  <GameDraggablePanel :panelId="'gm'">
    <div class="game-panel" v-if="showPanel">
      <div class="game-panel-content">
        <h4>Game Master</h4>
        <h6>Create Game Object</h6>
        <div id="add-object-cont">
          <label>Game object path:</label>
          <input type="text" v-model="objectPath" placeholder="resource/gold">
          <label>Offset X:</label>
          <input type="number" v-model="offsetX" step="0.1" placeholder="0.0">
          <label>Offset Y:</label>
          <input type="number" v-model="offsetY" step="0.1" placeholder="0.0">
          <label>Additional properties:</label>
          <textarea placeholder='{"amount":100.0}' v-model="additionalProps"></textarea>
          <button type="button" class="rpgui-button gold" @click="createGameObject"><p>Create</p></button>
        </div>
        <button type="button" class="rpgui-button" @click="showPanel = false"><p>Close</p></button>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      objectPath: "",
      offsetX: "0.0",
      offsetY: "0.0",
      additionalProps: "",
      showPanel: false,
    }
  },

  created() {
    EventBus.$on("show-gm-panel", this.showPanelWindow)
  },

  beforeDestroy() {
    EventBus.$off("show-gm-panel", this.showPanelWindow)
  },

  methods: {
    showPanelWindow() {
      this.showPanel = true
    },
    createGameObject() {
      EventBus.$emit("perform-game-action", {
        cmd: "gm_create_object",
        params: {
          "object_path": this.objectPath,
          "offset_x": parseFloat(this.offsetX),
          "offset_y": parseFloat(this.offsetY),
          "additional_props": this.additionalProps,
        }
      })
    }
  }
}
</script>


<style lang="scss">
#add-object-cont {
  width: 300px;
}
h6 {
  color: white;
}
</style>
