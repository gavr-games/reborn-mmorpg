<template>
  <div class="game-item">
    <div @contextmenu="showActions($event)">
      <div class="item-amount" v-if="item['amount'] > 0">{{ item['amount'] }}</div>
      <GameItemsIcon v-bind:item="item['kind']" />
    </div>
    <div class="actions-menu game-panel" v-if="showActionsMenu">
      <div class="game-panel-content">
        <div v-for="(action, actionKey) in item.actions" :key="actionKey" class="action-item" @click="handleAction(actionKey)">
          {{ actionKey }}
        </div>
        <div class="action-item" @click="showActionsMenu = false">
          close
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  props: ["item"],

  data() {
    return {
      showActionsMenu: false,
    }
  },

  methods: {
    showActions(e) {
      e.preventDefault()
      this.showActionsMenu = true
    },
    handleAction(actionKey) {
      this.showActionsMenu = false
      EventBus.$emit("perform-game-action", {
        cmd: this.item.actions[actionKey].cmd,
        params: this.item.actions[actionKey].params.replace("self", this.item.id)
      });
    },
  }
}
</script>


<style lang="scss">
.game-item {
  width: 32px;
  height: 32px;
  .actions-menu {
    position: absolute;
    margin-left: 10px;
    z-index: 999;
    .action-item {
      padding-bottom: 5px;
      color: white;
      &:hover {
        text-decoration: underline;
      }
    }
  }
  .item-amount {
    color: white;
    background-color: black;
    position: absolute;
    font-size: 5px;
  }
}
</style>
