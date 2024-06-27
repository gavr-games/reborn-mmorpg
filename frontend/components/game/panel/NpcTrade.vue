<template>
  <GameDraggablePanel :panel-id="'npc_trade'">
    <div v-if="showNpcTradePanel" id="npc_trade-panel" class="game-panel">
      <GameCloseIcon :close-callback="close" />
      <div class="game-panel-content">
        <h4 class="heading" @click="toggleExpandTab('sells')">
          Buy
        </h4>
        <div v-if="expandTabs['sells']" class="items">
          <div v-for="(item, itemName) in sellItems" :key="itemName" class="item">
            <GameItemsIcon :item="itemName.split('/')[1]" />:{{ item.amount }}
            for
            <GameItemsIcon :item="item.resource" />: {{ item.price }}
            <button type="button" class="rpgui-button" @click="buyItem(itemName)">
              <p>Buy</p>
            </button>
          </div>
        </div>
        <h4 class="heading" @click="toggleExpandTab('buys')">
          Sell
        </h4>
        <div v-if="expandTabs['buys']" class="items">
          <div v-for="(item, itemName) in buyItems" :key="itemName" class="item">
            <GameItemsIcon :item="itemName.split('/')[1]" />:{{ item.amount }}
            for
            <GameItemsIcon :item="item.resource" />: {{ item.price }}
            <button type="button" class="rpgui-button" @click="sellItem(itemName)">
              <p>Sell</p>
            </button>
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
      showNpcTradePanel: false,
      npcInfo: {},
      sellItems: {},
      buyItems: {},
      expandTabs: {}
    }
  },

  created () {
    EventBus.$on('npc_trade_info', this.showNpcTradeInfo)
  },

  beforeDestroy () {
    EventBus.$off('npc_trade_info', this.showNpcTradeInfo)
  },

  methods: {
    showNpcTradeInfo (data) {
      this.showNpcTradePanel = true
      this.npcInfo = data
      this.sellItems = data.sells
      this.buyItems = data.buys
    },
    toggleExpandTab (skillName) {
      if (this.expandTabs[skillName]) {
        this.expandTabs[skillName] = false
      } else {
        this.expandTabs[skillName] = true
      }
      this.$forceUpdate()
    },
    buyItem (itemName) {
      EventBus.$emit('perform-game-action', {
        cmd: 'npc_buy_item',
        params: {
          npc_id: this.npcInfo.id,
          item_name: itemName,
          amount: 1
        }
      })
    },
    sellItem (itemName) {
      EventBus.$emit('perform-game-action', {
        cmd: 'npc_sell_item',
        params: {
          npc_id: this.npcInfo.id,
          item_name: itemName,
          amount: 1
        }
      })
    },
    close () {
      this.showNpcTradePanel = false
    }
  }
}
</script>

<style lang="scss">
#npc_trade-panel {
  color: white;
  .heading {
    margin: 0px;
    padding-bottom: 8px;
  }
  .item {
    color: white;
    border: 1px solid white;
    padding: 5px;
  }
}
</style>
