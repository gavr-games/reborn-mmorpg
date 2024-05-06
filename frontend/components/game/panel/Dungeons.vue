<template>
  <GameDraggablePanel :panelId="'dungeons'">
    <div v-if="showDungeonsPanel" id="dungeons-panel" class="game-panel">
      <div class="game-panel-content">
        <h4 class="heading">
          Dungeons
        </h4>
        <div v-if="currentDungeonId === null">
          <label>Level:</label>
          <input
            v-model="dungeonLevel"
            type="number"
            step="1"
            placeholder="1"
            min="1"
            :max="maxDungeonLevel"
          >
          <label>Select Dragons (Max 3):</label>
          <div v-for="dragon in dragons" :key="dragon.id">
            <div :class="`dragon ${isDragonSelected(dragon.id) ? 'selected' : ''}`">
              <div class="dragon-title" @click="toggleExpandDragon(dragon.id)">
                {{ dragon.kind }} Lvl. {{ dragon.level }}
              </div>
              <div v-if="expandDragons[dragon.id]">
                <p class="dragon-description">
                  {{ dragon.alive ? 'Alive' : 'Dead' }} | Exp. {{ dragon.experience }}
                </p>
                <div v-if="dragon.alive">
                  <button v-if="!isDragonSelected(dragon.id)" type="button" class="rpgui-button" @click="selectDragon(dragon.id)">
                    <p>Select</p>
                  </button>
                  <button v-if="isDragonSelected(dragon.id)" type="button" class="rpgui-button" @click="deselectDragon(dragon.id)">
                    <p>Deselect</p>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-if="currentDungeonId !== null">
          <p>You have already started a dungeon.</p>
        </div>
        <button type="button" class="rpgui-button" @click="goToDungeon()">
          <p>Go To Dungeon</p>
        </button>
        <button type="button" class="rpgui-button" @click="showDungeonsPanel = false">
          <p>Close</p>
        </button>
      </div>
    </div>
  </GameDraggablePanel>
</template>

<script>
import { EventBus } from '~/plugins/game/event_bus'

const MAX_DRAGONS = 3

export default {
  data () {
    return {
      showDungeonsPanel: false,
      dragons: [],
      maxDungeonLevel: 1,
      currentDungeonId: null,
      dungeonLevel: 1,
      selectedDragons: [],
      expandDragons: {}
    }
  },

  created () {
    EventBus.$on('dungeons_info', this.showDungeonsInfo)
  },

  beforeDestroy () {
    EventBus.$off('dungeons_info', this.showDungeonsInfo)
  },

  methods: {
    showDungeonsInfo (data) {
      this.showDungeonsPanel = true
      this.maxDungeonLevel = data.max_dungeon_lvl
      this.currentDungeonId = data.current_dungeon_id
      this.dragons = data.dragons
    },
    toggleExpandDragon (dragonId) {
      if (this.expandDragons[dragonId]) {
        this.expandDragons[dragonId] = false
      } else {
        this.expandDragons[dragonId] = true
      }
      this.$forceUpdate()
    },
    selectDragon (dragonId) {
      if (this.selectedDragons.length < MAX_DRAGONS) {
        this.selectedDragons.push(dragonId)
        this.$forceUpdate()
      }
    },
    deselectDragon (dragonId) {
      const dragonIndex = this.selectedDragons.indexOf(dragonId)
      if (dragonIndex !== -1) {
        this.selectedDragons.splice(dragonIndex, 1)
        this.$forceUpdate()
      }
    },
    isDragonSelected (dragonId) {
      return this.selectedDragons.includes(dragonId)
    },
    goToDungeon () {
      EventBus.$emit('perform-game-action', {
        cmd: 'go_to_dungeon',
        params: {
          level: this.dungeonLevel,
          dragons: this.selectedDragons
        }
      })
    }
  }
}
</script>

<style lang="scss">
#dungeons-panel {
  color: white;
  .heading {
    margin-top: 0px;
  }
  .dragon {
    &.selected {
      background-color: burlywood;
    }
    .dragon-title {
      color: white;
      &:hover {
        text-decoration: underline;
      }
    }
    border: 1px solid white;
    padding: 5px;
    .dragon-description {
      color: grey;
      margin: 0px;
      font-size: 10px;
    }
    .actions {
      display: flex;
    }
  }
}
</style>
