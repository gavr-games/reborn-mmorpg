<template>
  <div id="npc_trade-panel" class="game-panel" v-if="showNpcTradePanel">
    <div class="game-panel-content">
      <h4 class="heading" @click="toggleExpandTab('sells')">Buy</h4>
      <div class="items" v-if="expandTabs['sells']">
        <div v-for="(item, itemName) in sellItems" :key="itemName" class="item">
          <GameItemsIcon v-bind:item="itemName.split('/')[1]" />:{{ item.amount }}
          for
          <GameItemsIcon v-bind:item="item.resource" />: {{ item.price }}
          <button type="button" class="rpgui-button" @click="buyItem(itemName)"><p>Buy</p></button>
        </div>
      </div>
      <h4 class="heading" @click="toggleExpandTab('buys')">Sell</h4>
      <div class="items" v-if="expandTabs['buys']">
      </div>
      <button type="button" class="rpgui-button" @click="showNpcTradePanel = false"><p>Close</p></button>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showNpcTradePanel: false,
      npcInfo: {},
      sellItems: {},
      buyItems: {},
      expandTabs: {},
    }
  },

  created() {
    EventBus.$on("npc_trade_info", this.showNpcTradeInfo)
  },

  beforeDestroy() {
    EventBus.$off("npc_trade_info", this.showNpcTradeInfo)
  },

  methods: {
    showNpcTradeInfo(data) {
      this.showNpcTradePanel = true
      this.npcInfo = data
      this.sellItems = data.sells
    },
    toggleExpandTab(skillName) {
      if (this.expandTabs[skillName]) {
        this.expandTabs[skillName] = false
      } else {
        this.expandTabs[skillName] = true
      }
      this.$forceUpdate()
    },
    buyItem(itemName) {
      EventBus.$emit("perform-game-action", {
        cmd: "npc_buy_item",
        params: {
          "npc_id": this.npcInfo.id,
          "item_name": itemName,
          "amount": 1,
        }
      });
    }
  }
}
</script>


<style lang="scss">
#npc_trade-panel {
  position: absolute;
  top: 50px;
  left: 650px;
  color: white;
  .heading {
    margin-top: 0px;
  }
  .item {
    color: white;
    border: 1px solid white;
    padding: 5px;
  }
}
</style>
