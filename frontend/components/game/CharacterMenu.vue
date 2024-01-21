<template>
  <div id="character-menu" class="game-panel">
    <div class="game-panel-content">
      <div class="menu-item hero-icon" @click="getCharacterInfo" title="Character"></div>
      <div class="menu-item craft-icon" @click="getCraftAtlas" title="Craft"></div>
      <div class="menu-item map-icon"  @click="showMap" title="Map"></div>
      <div class="menu-item town-icon"  @click="townTeleport" title="Teleport to Town"></div>
      <div class="menu-item obelisk-icon"  @click="claimTeleport" title="Teleport to Claim"></div>
      <div class="menu-item gm-icon"  @click="showGMPanel" title="Game Master" v-if="showGMIcon"></div>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showGMIcon: false,
    }
  },

  created() {
    EventBus.$on("my-character-info", this.showGameMasterIcon)
  },

  beforeDestroy() {
    EventBus.$off("my-character-info", this.showGameMasterIcon)
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
    townTeleport() {
      EventBus.$emit("perform-game-action", {
        cmd: "town_teleport",
        params: {}
      });
    },
    claimTeleport() {
      EventBus.$emit("perform-game-action", {
        cmd: "claim_teleport",
        params: {}
      });
    },
    showGameMasterIcon(characterData) {
      if (characterData.Properties["game_master"] === true) {
        this.showGMIcon = true
      }
    },
    showGMPanel() {
      EventBus.$emit("show-gm-panel", {});
    }
  }
}
</script>


<style lang="scss">
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
  .town-icon {
    background-image: url("~assets/img/icons/town.png");
    background-repeat: no-repeat;
  }
  .obelisk-icon {
    background-image: url("~assets/img/icons/obelisk.png");
    background-repeat: no-repeat;
  }
  .gm-icon {
    background-image: url("~assets/img/icons/gm.png");
    background-repeat: no-repeat;
  }
}
</style>
