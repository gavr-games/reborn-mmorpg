<template>
  <div id="experience-bar" class="rpgui-progress" data-rpguitype="progress">
    <div class="rpgui-progress-track">
      <div class="exp-text">
        Lvl. {{ level }} Exp. {{ experience }} / {{ expNextLevel }}
      </div>
      <div class="rpgui-progress-fill" :style="`left: 0px; width: ${progress}%;`" />
    </div>
    <div class="rpgui-progress-left-edge" />
    <div class="rpgui-progress-right-edge" />
  </div>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

const EXP_AMOUNT_DIVIDER = 0.05
const EXP_INCREASE_POWER = 2.0

export default {
  data () {
    return {
      experience: 0,
      level: 0,
      currentCharId: 0
    }
  },

  computed: {
    expNextLevel () {
      return Math.pow((this.level + 1) / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER) - Math.pow(this.level / EXP_AMOUNT_DIVIDER, EXP_INCREASE_POWER)
    },
    progress () {
      return Math.ceil(this.experience * 100 / this.expNextLevel)
    }
  },

  created () {
    EventBus.$on('my-character-info', this.setExperienceAndLevel)
    EventBus.$on('set_exp', this.setExperience)
    EventBus.$on('set_level', this.setLevel)
  },

  beforeDestroy () {
    EventBus.$off('my-character-info', this.setExperienceAndLevel)
    EventBus.$off('set_exp', this.setExperience)
    EventBus.$off('set_level', this.setLevel)
  },

  methods: {
    setExperienceAndLevel (characterData) {
      this.experience = characterData.Properties.experience
      this.level = characterData.Properties.level
      this.currentCharId = characterData.Id
    },
    setExperience (data) {
      if (data.object_id === this.currentCharId) {
        this.experience = data.exp
      }
    },
    setLevel (data) {
      if (data.object_id === this.currentCharId) {
        this.level = data.level
        // TODO: add some levelup animation
      }
    }
  }
}
</script>

<style lang="scss">
#experience-bar {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 30%;
  .exp-text {
    position: absolute;
    color: white;
    z-index: 1;
    margin-top: 16px;
    margin-left: 5px;
    font-size: 8px;
  }
}
</style>
