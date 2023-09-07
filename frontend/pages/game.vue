<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="rpgui-container framed-golden">0</div>
    <div id="character-menu" class="rpgui-container framed-golden">
      <div class="rpgui-icon helmet-slot" @click="getCharacterInfo"></div>
    </div>
    <div id="character-info-panel" class="rpgui-container framed-golden" v-if="showCharacterInfoPanel">
      <h4>Character</h4>
      <div v-for="(slotItem, slotKey) in characterInfo.slots" :key="slotKey">
        <span>{{ slotKey }}:</span> <span v-if="slotItem"><i :class="`game-item-icon ${slotItem['kind']}`" :title="slotItem['kind']"></i></span>
      </div>
      <button type="button" class="rpgui-button" @click="showCharacterInfoPanel = false"><p>Close</p></button>
    </div>
    <div id="chat" class="rpgui-container framed-golden">
      <div id="chat-messages" class="rpgui-container framed-grey">
        <div class="message" v-for="index in chatMessages.length" :key="'msg'+index">
          {{ chatMessages[index - 1] }}
        </div>
      </div>
      <input id="input-chat-message" placeholder="type your message here" v-model="chatMessage" v-on:keyup.enter="sendChatMessage" autocomplete="off">
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      chatMessage: '',
      chatMessages: [],
      showCharacterInfoPanel: false,
      characterInfo: {
        "slots": {}
      },
    }
  },

  mounted() {
    if (!this.$auth.loggedIn || !this.$store.state.characters.selectedCharacterId) {
      this.$router.push('login')
    } else {
      this.$game.init(this.$auth.strategy.token.get(), this.$store.state.characters.selectedCharacterId)
    }
  },
  created() {
    EventBus.$on("new-chat-message", this.addNewChatMessage)
    EventBus.$on("character_info", this.showCharacterInfo)
  },
  beforeDestroy() {
    this.$game.destroy()
    EventBus.$off("new-chat-message")
    EventBus.$off("character_info")
  },

  computed: {
    charId() {
      return this.$store.state.characters.selectedCharacterId
    }
  },

  methods: {
    sendChatMessage() {
      EventBus.$emit("send-chat-message", this.chatMessage);
      this.chatMessage = ''
    },
    addNewChatMessage(message) {
      this.chatMessages.push(message)
      const container = this.$el.querySelector("#chat-messages");
      setTimeout(() => {
        container.scrollTo(0, container.scrollHeight);
      }, 100);
    },
    getCharacterInfo() {
      EventBus.$emit("get-character-info", {});
    },
    showCharacterInfo(data) {
      this.showCharacterInfoPanel = true
      this.characterInfo = data
    },
  }
}
</script>
