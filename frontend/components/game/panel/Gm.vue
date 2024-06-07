<template>
  <GameDraggablePanel :panelId="'gm'">
    <div v-if="showPanel" class="game-panel">
      <div class="game-panel-content">
        <h4>Game Master</h4>
        <a @click="showCreateGameObjectForm">Create Game Object</a> <br>
        <a @click="showUpdateGameObjectPropertiesForm">Update Game Object Properties</a>
        <div v-if="showCreateGameObject" id="add-object-cont">
          <h6>Create Game Object</h6>
          <label>Game object path:</label>
          <input v-model="objectPath" type="text" placeholder="resource/gold">
          <label>Offset X:</label>
          <input v-model="offsetX" type="number" step="0.1" placeholder="0.0">
          <label>Offset Y:</label>
          <input v-model="offsetY" type="number" step="0.1" placeholder="0.0">
          <label>Additional properties:</label>
          <textarea v-model="additionalProps" placeholder='{"amount":100.0}'/>
          <button type="button" class="rpgui-button gold" @click="createGameObject">
            <p>Create</p>
          </button>
        </div>
        <div v-if="showUpdateGameObjectProperties" id="update-properties-cont">
          <h6>Update Game Object Properties</h6>
          <label>Object Id:</label>
          <input v-model="objectId" type="text" placeholder="3bf3d937-282c-4410-a2dc-2d0ca207e386">
          <label>Update Properties:</label>
          <textarea v-model="updateProps" placeholder='{"amount":100.0}'/>
          <button type="button" class="rpgui-button gold" @click="updateGameObjectProperties">
            <p>Update</p>
          </button>
        </div>
        <button type="button" class="rpgui-button" @click="showPanel = false">
          <p>Close</p>
        </button>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      objectPath: '',
      offsetX: '0.0',
      offsetY: '0.0',
      additionalProps: '',
      objectId: '',
      updateProps: '',
      showPanel: false,
      showCreateGameObject: true,
      showUpdateGameObjectProperties: false
    }
  },

  created () {
    EventBus.$on('show-gm-panel', this.showPanelWindow)
  },

  beforeDestroy () {
    EventBus.$off('show-gm-panel', this.showPanelWindow)
  },

  methods: {
    showPanelWindow () {
      this.showPanel = true
    },
    showCreateGameObjectForm () {
      this.showCreateGameObject = true
      this.showUpdateGameObjectProperties = false
    },
    showUpdateGameObjectPropertiesForm () {
      this.showCreateGameObject = false
      this.showUpdateGameObjectProperties = true
    },
    createGameObject () {
      EventBus.$emit('perform-game-action', {
        cmd: 'gm_create_object',
        params: {
          object_path: this.objectPath,
          offset_x: parseFloat(this.offsetX),
          offset_y: parseFloat(this.offsetY),
          additional_props: this.additionalProps
        }
      })
    },
    updateGameObjectProperties () {
      EventBus.$emit('perform-game-action', {
        cmd: 'gm_update_properties',
        params: {
          object_id: this.objectId,
          update_props: this.updateProps
        }
      })
    }
  }
}
</script>

<style lang="scss">
#add-object-cont, #update-properties-cont {
  width: 300px;
}
h6 {
  color: white;
}
</style>
