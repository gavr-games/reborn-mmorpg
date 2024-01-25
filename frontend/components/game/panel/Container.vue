<template>
  <GameDraggablePanel v-bind:panelId="container.id">
    <div :class="`game-container size-${container.size}`" v-if="showContainerPanel">
      <a href="#" class="close-btn" @click="close()"></a>
      <div v-for="(item, key) in container.items" :key="key" class="slot">
        <div
          draggable
          @dragstart="startDrag($event, item)"
          @dragend="endDrag($event)"
        >
          <GameItem v-if="item !== null" v-bind:item="item" />
        </div>
        <div class="empty-slot" @dragover="allowDrag" @drop="onDrop($event, key)" v-if="item === null" />
      </div>
    </div>
  </GameDraggablePanel>
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
    EventBus.$on("update-item-in-container", this.updateItem)
  },

  mounted() {
    let openedContainers = localStorage.getItem("opened_containers")
    if (openedContainers) {
      openedContainers = JSON.parse(openedContainers)
    } else {
      openedContainers = []
    }
    if (!openedContainers.includes(this.container.id)) {
      openedContainers.push(this.container.id)
      localStorage.setItem("opened_containers", JSON.stringify(openedContainers))
    }
  },

  beforeDestroy() {
    EventBus.$off("put_item_to_container", this.addItem)
    EventBus.$off("remove_item_from_container", this.removeItem)
    EventBus.$off("update-item-in-container", this.updateItem)
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
    updateItem(data) {
      if (data.container_id === this.container.id) {
        this.container.items.forEach((item, index) => {
          if (item && item.id === data.item.Id) {
            this.container.items[index] = data.item.Properties
            this.$forceUpdate()
            return
          }
        })
      }
    },
    startDrag(evt, item) {
      if (item !== null) {
        evt.dataTransfer.dropEffect = 'move'
        evt.dataTransfer.effectAllowed = 'move'
        evt.dataTransfer.setData('item_id', item.id)
        evt.stopPropagation()
      }
    },
    endDrag(evt, item) {
      evt.stopPropagation()
    },
    onDrop(evt, pos) {
      evt.stopPropagation()
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
    },
    close() {
      this.showContainerPanel = false
      let openedContainers = localStorage.getItem("opened_containers")
      if (openedContainers) {
        openedContainers = JSON.parse(openedContainers)
        const index = openedContainers.indexOf(this.container.id);
        if (index !== -1) {
          openedContainers.splice(index, 1);
          localStorage.setItem("opened_containers", JSON.stringify(openedContainers))
        }
      }
    }
  }
}
</script>


<style lang="scss">
.game-container {
  &.size-8 {
    width: 330px;
    height: 165px;
    padding-top: 14px;
    padding-left: 14px;
    background-image: url("~assets/img/container-8.png");
  }
  &.size-4 {
    width: 165px;
    height: 165px;
    padding-top: 14px;
    padding-left: 14px;
    background-image: url("~assets/img/backpack-16.png");
  }
  &.size-2 {
    width: 84px;
    height: 87px;
    padding-top: 14px;
    padding-left: 14px;
    background-image: url("~assets/img/container-4.png");
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
