<template>
  <div id="character-menu" class="rpgui-container framed-golden">
    <div class="rpgui-icon helmet-slot" @click="getCharacterInfo"></div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
    }
  },

  methods: {
    getCharacterInfo() {
      EventBus.$emit("get-character-info", {});
    },
  }
}
</script>


<style>
#character-menu {
  position: absolute;
  top: 0;
  right: 0;
  .rpgui-icon {
    width: 32px;
    height: 32px;
  }
}
</style>
