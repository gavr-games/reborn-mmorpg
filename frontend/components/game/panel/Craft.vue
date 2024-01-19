<template>
  <GameDraggablePanel :panelId="'craft'">
    <div id="craft-panel" class="game-panel" v-if="showCraftPanel">
      <div class="game-panel-content">
        <h4 class="heading">Craft</h4>
        <div v-for="(craftItems, skillName) in craftInfo" :key="skillName">
          <div class="skill">
            <div class="skill-title" @click="toggleExpandSkill(skillName)">{{ skillName }}</div>
            <div class="craft-items" v-if="expandSkills[skillName]">
              <div class="craft-item" v-for="(item, itemKey) in craftItems" :key="itemKey">
                <span>{{ item["title"] }}</span>
                <p class="item-description">{{ item["description"] }}</p>
                <div class="craft-resources">
                  <p v-for="(resourceCount, resourceName) in item.resources" :key="resourceName">
                    <GameItemsIcon v-bind:item="resourceName" />:{{ resourceCount }}
                  </p>
                  <button type="button" class="rpgui-button" @click="craftItem(itemKey, item)"><p>Craft</p></button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <button type="button" class="rpgui-button" @click="showCraftPanel = false"><p>Close</p></button>
      </div>
    </div>
  </GameDraggablePanel>
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
      if (item.place_in_real_world) {
        EventBus.$emit("select-coords-and-rotation", {
          "item_key": itemKey,
          "item": item
        });
      } else {
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


<style lang="scss">
#craft-panel {
  color: white;
  .heading {
    margin-top: 0px;
  }
  .skill {
    .skill-title {
      color: white;
      &:hover {
        text-decoration: underline;
      }
    }
    .craft-item {
      border: 1px solid white;
      padding: 5px;
      .item-description {
        color: grey;
        margin: 0px;
        font-size: 10px;
      }
      .craft-resources {
        display: flex;
      }
    }
  }
}
</style>
