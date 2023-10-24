<template>
  <div :class="`rpgui-container framed-golden game-container size-${container.size}`" v-if="showContainerPanel">
    <div v-for="(item, key) in container.items" :key="key" class="slot">
      <GameItem v-if="item !== null" v-bind:item="item" />
      <div class="empty-slot" v-if="item === null" />
    </div>
    <button type="button" class="rpgui-button" @click="showContainerPanel = false"><p>Close</p></button>
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
    width: 200px;
  }
  .slot {
    display: inline-block;
    border: 1px solid black;
    margin-right: 2px;
  }
  .empty-slot {
    width: 32px;
    height: 32px;
  }
}
</style>
