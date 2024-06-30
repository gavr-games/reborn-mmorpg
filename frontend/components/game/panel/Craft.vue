<template>
  <GameDraggablePanel :panel-id="'craft'">
    <div v-if="showCraftPanel" id="craft-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4 class="heading">
          Craft
        </h4>
        <div id="craft-items-cont">
          <div id="searchable-list">
            <input v-model="searchTerm" type="text" placeholder="Search" @keyup="search">
            <div v-if="searchTerm === ''">
              <div v-for="(craftItems, skillName) in craftInfo" :key="skillName">
                <div class="skill">
                  <div class="skill-title" @click="toggleExpandSkill(skillName)">
                    {{ capitalizeFirstLetter(skillName) }}
                  </div>
                  <div v-if="expandSkills[skillName]" class="craft-items">
                    <div v-for="(item, itemKey) in craftItems" :key="itemKey">
                      <a href="#" class="craft-item-link" @click="selectItem(itemKey)">{{ item["title"] }}</a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div v-if="searchResults">
              <div v-for="(item, itemKey) in searchResults" :key="itemKey">
                <a href="#" class="craft-item-link" @click="selectItem(itemKey)">{{ item["title"] }}</a>
              </div>
            </div>
          </div>
          <div v-if="selectedItem" id="selected-item">
            <div class="craft-item">
              <div class="item-title">
                <GameItemsIcon :item="selectedItemKey" />
                <p>{{ selectedItem["title"] }}</p>
              </div>
              <p class="item-description">
                {{ selectedItem["description"] }}
              </p>
              <div class="craft-resources">
                <div v-for="(resourceCount, resourceName) in selectedItem.resources" :key="resourceName" @click="selectItem(resourceName)">
                  <GameItemsIcon :item="resourceName" />{{ resourceCount }}
                </div>
              </div>
              <button type="button" class="rpgui-button" @click="craftItem(selectedItemKey, selectedItem)">
                <p>Craft</p>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

export default {
  data () {
    return {
      showCraftPanel: false,
      searchTerm: '',
      searchResults: {},
      craftInfo: {},
      craftItemsInfo: {},
      expandSkills: {},
      selectedItem: null,
      selectedItemKey: ''
    }
  },

  created () {
    EventBus.$on('craft_atlas', this.showCraftInfo)
  },

  beforeDestroy () {
    EventBus.$off('craft_atlas', this.showCraftInfo)
  },

  methods: {
    search () {
      this.searchResults = null
      if (this.searchTerm !== '') {
        for (const key in this.craftItemsInfo) {
          if (key.toLowerCase().includes(this.searchTerm.toLowerCase())) {
            if (this.searchResults === null) {
              this.searchResults = {}
            }
            this.searchResults[key] = this.craftItemsInfo[key]
          }
        }
      }
    },
    capitalizeFirstLetter (string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    showCraftInfo (data) {
      this.showCraftPanel = true
      this.craftItemsInfo = data
      const skillBasedData = {}
      Object.entries(data).forEach((entry) => {
        const [itemName, item] = entry
        if (!skillBasedData[item.skill]) {
          skillBasedData[item.skill] = {}
        }
        skillBasedData[item.skill][itemName] = item
      })
      this.craftInfo = skillBasedData
    },
    toggleExpandSkill (skillName) {
      if (this.expandSkills[skillName]) {
        this.expandSkills[skillName] = false
      } else {
        this.expandSkills[skillName] = true
      }
      this.$forceUpdate()
    },
    selectItem (itemKey) {
      if (itemKey in this.craftItemsInfo) {
        this.selectedItemKey = itemKey
        this.selectedItem = this.craftItemsInfo[itemKey]
      }
    },
    craftItem (itemKey, item) {
      if (item.place_in_real_world) {
        this.showCraftPanel = false
        EventBus.$emit('select-coords-and-rotation', {
          item_key: itemKey,
          item,
          cmd: 'craft',
          callback: (x, y, rotation) => {
            EventBus.$emit('perform-game-action', {
              cmd: 'craft',
              params: {
                item_name: itemKey,
                inputs: {
                  coordinates: {
                    x,
                    y
                  },
                  rotation
                }
              }
            })
          }
        })
      } else {
        EventBus.$emit('perform-game-action', {
          cmd: 'craft',
          params: {
            item_name: itemKey,
            inputs: []
          }
        })
      }
    },
    close () {
      this.showCraftPanel = false
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
  #craft-items-cont {
    display: flex;
  }
  #searchable-list {
    border-right: 1px solid #ccc;
    padding-right: 5px;
    input {
      line-height: 10px;
      margin-bottom: 7px;
    }
  }
  #selected-item {
    margin-left: 20px;
    .craft-item {
      .item-title {
        display: flex;
        p {
          margin-left: 10px;
          line-height: 10px;
        }
      }
      .item-description {
        color: grey;
        margin: 0px;
        font-size: 10px;
        margin-bottom: 10px;
      }
      .craft-resources {
        display: flex;
        margin-bottom: 10px;
      }
    }
  }
  .skill {
    .skill-title {
      color: white;
      &:hover {
        text-decoration: underline;
      }
    }
  }
  .craft-item-link {
    display: block;
    border: none;
    color: #ccc;
    font-size: 10px;
    line-height: 15px;
    padding-left: 10px;
  }
}
</style>
