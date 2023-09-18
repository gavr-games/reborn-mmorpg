<template>
  <div id="craft-panel" class="rpgui-container framed-golden" v-if="showCraftPanel">
    <h4>Craft</h4>
    <div v-for="(craftItems, skillName) in craftInfo" :key="skillName">
      <div class="skill">
        <div class="skill-title" @click="toggleExpandSkill(skillName)">{{ skillName }}</div>
        <div class="craft-items" v-if="expandSkills[skillName]">
          <div class="craft-item" v-for="(item, itemKey) in craftItems" :key="itemKey">
            <span>{{ item["title"] }}</span>
            <p>{{ item["description"] }}</p>
            <p v-for="(resourceCount, resourceName) in item.resources" :key="resourceName">
              <GameItemsIcon v-bind:item="resourceName" />: {{ resourceCount }}
            </p>
            <button type="button" class="rpgui-button" @click="craftItem(itemKey, item)"><p>Craft</p></button>
          </div>
        </div>
      </div>
    </div>
    <button type="button" class="rpgui-button" @click="showCraftPanel = false"><p>Close</p></button>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showCraftPanel: false,
      craftInfo: {},
      expandSkills: {},
    }
  },

  created() {
    EventBus.$on("craft_atlas", this.showCraftInfo)
  },

  beforeDestroy() {
    EventBus.$off("craft_atlas", this.showCraftInfo)
  },

  methods: {
    showCraftInfo(data) {
      this.showCraftPanel = true
      const skillBasedData = {}
      Object.entries(data).forEach(entry => {
        const [itemName, item] = entry;
        if (!skillBasedData[item.skill]) {
          skillBasedData[item.skill] = {}
        }
        skillBasedData[item.skill][itemName] = item
      });
      this.craftInfo = skillBasedData
    },
    toggleExpandSkill(skillName) {
      if (this.expandSkills[skillName]) {
        this.expandSkills[skillName] = false
      } else {
        this.expandSkills[skillName] = true
      }
      this.$forceUpdate()
    },
    craftItem(itemKey, item) {
      if (item.inputs.length === 0) {
        EventBus.$emit("perform-game-action", {
          cmd: "craft",
          params: {
            "item_name": itemKey,
            "inputs": []
          }
        });
      }
    }
  }
}
</script>


<style>
#craft-panel {
  position: absolute;
  top: 50px;
  left: 50px;
  .skill {
    .skill-title {
      &:hover {
        color: white;
      }
    }
    .craft-item {
      border: 1px solid white;
      padding: 5px;
    }
  }
}
</style>
