<template>
  <div id="character-menu" class="rpgui-container framed-golden">
    <div class="rpgui-icon menu-item helmet-slot" @click="getCharacterInfo"></div>
    <div class="rpgui-icon menu-item potion-slot" @click="getCraftAtlas"></div>
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
      EventBus.$emit("perform-game-action", {
        cmd: "get_character_info",
        params: {}
      });
    },
    getCraftAtlas() {
      EventBus.$emit("perform-game-action", {
        cmd: "get_craft_atlas",
        params: {}
      });
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
  .menu-item {
    display: block;
    margin-bottom: 5px;
  }
}
</style>
