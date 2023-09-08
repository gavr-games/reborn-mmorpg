<template>
  <div class="game-item">
    <i :class="`game-item-icon ${item['kind']}`" :title="item['kind']" @contextmenu="showActions($event)"></i>
    <div class="actions-menu" v-if="showActionsMenu">
      <div v-for="(action, actionKey) in item.actions" :key="actionKey" class="action-item" @click="handleAction(actionKey)">
        {{ actionKey }}
      </div>
      <div class="action-item" @click="showActionsMenu = false">
        close
      </div>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  props: ['item'],

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


<style>
.game-item {
  .game-item-icon {
    display: inline-block;
    width: 32px;
    height: 32px;
    background-image: url("~assets/img/icons/game-items-icons.png");
    &.backpack {
      background-position: 224px 608px;
    }
  }
  .actions-menu {
    position: absolute;
    margin-left: 10px;
    padding: 10px;
    border: 2px solid black;
    background-color: #884c34;
    .action-item {
      margin-bottom: 5px;
      &:hover {
        color: white;
      }
    }
  }
}
</style>
