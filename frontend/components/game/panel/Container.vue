<template>
  <div :class="`game-container size-${container.size}`" v-if="showContainerPanel">
    <a href="#" class="close-btn" @click="showContainerPanel = false"></a>
    <div v-for="(item, key) in container.items" :key="key" class="slot">
      <GameItem v-if="item !== null" v-bind:item="item" />
      <div class="empty-slot" v-if="item === null" />
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  props: ["container"],
  data() {
    return {
      showContainerPanel: true,
    }
  },

  created() {
    EventBus.$on("put_item_to_container", this.addItem)
    EventBus.$on("remove_item_from_container", this.removeItem)
  },

  beforeDestroy() {
    EventBus.$off("put_item_to_container", this.addItem)
    EventBus.$off("remove_item_from_container", this.removeItem)
  },

  methods: {
    addItem(data) {
      if (data.container_id === this.container.id) {
        this.container.items[data.position] = data.item
        this.$forceUpdate()
      }
    },
    removeItem(data) {
      if (data.container_id === this.container.id) {
        this.container.items[data.position] = null
        this.$forceUpdate()
      }
    },
  }
}
</script>


<style>
.game-container {
  position: absolute;
  top: 50px;
  left: 250px;
  &.size-4 {
    width: 165px;
    height: 165px;
    padding-top: 14px;
    padding-left: 14px;
    background-image: url("~assets/img/backpack-16.png");
  }
  .slot {
    display: inline-block;
    width: 32px;
    height: 32px;
    margin-right: 8px;
    margin-top: 2px;
  }
  .empty-slot {
    width: 32px;
    height: 32px;
  }
  .close-btn {
    position: absolute;
    top: 0px;
    right: 0px;
    transform: rotate(45deg);
    width: 16px;
    height: 16px;
    background-image: url("~assets/img/close.png");
    &:hover {
      transform: rotate(0deg);
      transition: rotate 1s;
    }
  }
}
</style>
