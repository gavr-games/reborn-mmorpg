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

export default {
  props: ["container"],
  data() {
    return {
      showContainerPanel: true,
    }
  },
}
</script>


<style>
.game-container {
  position: absolute;
  top: 100px;
  left: 100px;
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
