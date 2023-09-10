<template>
    <div class="actions-menu" v-if="showActionsMenu" :style="`top: ${y}px; left: ${x}px;`">
      <div v-for="(action, actionKey) in item.Properties.actions"
        :key="actionKey"
        class="action-item"
        @click="handleAction(actionKey)">
        {{ actionKey }}
      </div>
      <div class="action-item" @click="showActionsMenu = false">
        close
      </div>
    </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      item: {
        actions: {},
      },
      x: 0,
      y: 0,
      showActionsMenu: false,
    }
  },

  created() {
    EventBus.$on("game-object-clicked", this.showActions)
  },

  beforeDestroy() {
    EventBus.$off("game-object-clicked", this.showActions)
  },

  methods: {
    showActions(data) {
      if (data.game_object.Properties.actions) {
        this.showActionsMenu = true
        this.item = data.game_object
        this.x = data.x
        this.y = data.y
      }
    },
    handleAction(actionKey) {
      this.showActionsMenu = false
      EventBus.$emit("perform-game-action", {
        cmd: this.item.Properties.actions[actionKey].cmd,
        params: this.item.Properties.actions[actionKey].params.replace("self", this.item.Properties.id)
      });
    },
  }
}
</script>


<style>
.actions-menu {
  position: absolute;
  padding: 10px;
  border: 2px solid black;
  background-color: #884c34;
  z-index: 999;
  .action-item {
    margin-bottom: 5px;
    &:hover {
      color: white;
    }
  }
}
</style>
