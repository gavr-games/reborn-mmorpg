<template>
  <div id="character-menu" class="game-panel">
    <div class="game-panel-content">
      <div class="menu-item hero-icon" @click="getCharacterInfo"></div>
      <div class="menu-item craft-icon" @click="getCraftAtlas"></div>
      <div class="menu-item map-icon"  @click="showMap"></div>
    </div>
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
    showMap() {
      EventBus.$emit("show-map", {});
    },
  }
}
</script>


<style>
#character-menu {
  position: absolute;
  top: 0;
  right: 0;
  .menu-item {
    display: block;
    width: 32px;
    height: 32px;
    padding-bottom: 5px;
    opacity: 0.8;
    &:hover {
      opacity: 1.0;
      cursor: url("~assets/img/cursor/point.png") 10 0, auto;
    }
  }
  .map-icon {
    background-image: url("~assets/img/icons/map.png");
    background-repeat: no-repeat;
  }
  .hero-icon {
    background-image: url("~assets/img/icons/hero.png");
    background-repeat: no-repeat;
  }
  .craft-icon {
    background-image: url("~assets/img/icons/craft.png");
    background-repeat: no-repeat;
  }
}
</style>
