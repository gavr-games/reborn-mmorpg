<template>
  <div :class="`game-container size-${container.size}`" v-if="showContainerPanel">
    <a href="#" class="close-btn" @click="showContainerPanel = false"></a>
    <div v-for="(item, key) in container.items" :key="key" class="slot">
      <div
        draggable
        @dragstart="startDrag($event, item)"
      >
        <GameItem v-if="item !== null" v-bind:item="item" />
      </div>
      <div class="empty-slot" @dragover="allowDrag" @drop="onDrop($event, key)" v-if="item === null" />
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
    startDrag(evt, item) {
      if (item !== null) {
        evt.dataTransfer.dropEffect = 'move'
        evt.dataTransfer.effectAllowed = 'move'
        evt.dataTransfer.setData('item_id', item.id)
      }
    },
    onDrop(evt, pos) {
      const itemID = evt.dataTransfer.getData('item_id')
      EventBus.$emit("perform-game-action", {
        cmd: "put_to_container",
        params: {
          "container_id": this.container.id,
          "position": pos,
          "item_id": itemID,
        }
      })
    },
    allowDrag(evt) {
      evt.preventDefault()
    }
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
    margin-top: 3px;
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
