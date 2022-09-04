<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <script src="https://preview.babylonjs.com/inspector/babylon.inspector.bundle.js"></script>
    <canvas id="game-canvas"></canvas>
    <div id="fps" class="rpgui-container framed-golden">0</div>
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
      chatMessages: []
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
  },
  beforeDestroy() {
    this.$game.destroy()
    EventBus.$off("new-chat-message")
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
    }
  }
}
</script>
