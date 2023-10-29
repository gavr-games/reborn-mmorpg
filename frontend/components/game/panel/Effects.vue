<template>
  <div id="effects-panel">
    <div v-for="(effect, effectId) in effects" :key="effectId" :class="`effect ${effect['group']}`" :title="effect['group']">
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      effects: {},
      currentPlayerId: 0,
    }
  },

  created() {
    this.currentPlayerId = this.$store.state.characters.selectedCharacterId
    EventBus.$on("add_object", this.showEffects)
    EventBus.$on("update_object", this.showEffects)
  },

  beforeDestroy() {
    EventBus.$off("add_object", this.showEffects)
    EventBus.$off("update_object", this.showEffects)
  },

  methods: {
    showEffects(data) {
      if (data.Properties && data.Properties.player_id && data.Properties.player_id == this.currentPlayerId) {
        this.effects = data.Effects
      }
    },
  }
}
</script>


<style>
#effects-panel {
  position: absolute;
  top: 170px;
  right: 3px;
  width: 32px;
  display: flex;
  flex-direction: column;
  .effect.potion_healing {
    width: 32px;
    height: 32px;
    margin-bottom: 2px;
    background-image: url("~assets/img/icons/potion_healing.png");
    background-size: 32px 32px;
  }
}
</style>
