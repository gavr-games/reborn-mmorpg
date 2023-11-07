<template>
  <div id="chat" class="game-panel">
    <div class="game-panel-content">
      <div id="chat-messages">
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
    }
  },

  created() {
    EventBus.$on("new-chat-message", this.addNewChatMessage)
    EventBus.$on("add_message", this.addMessage)
  },

  beforeDestroy() {
    EventBus.$off("new-chat-message", this.addNewChatMessage)
    EventBus.$off("add_message", this.addMessage)
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
    addMessage(data) {
      this.chatMessages.push(data.message)
      const container = this.$el.querySelector("#chat-messages");
      setTimeout(() => {
        container.scrollTo(0, container.scrollHeight);
      }, 100);
    }
  }
}
</script>


<style lang="scss">
#chat {
  width: 500px;
  height: 250px;
  position: absolute;
  bottom: 0;
  left: 0;

  #chat-messages {
    width: 100%;
    height: 210px;
    overflow-y: scroll;
    padding-top: 0;
    padding-bottom: 0;
    padding-left: 0;
    .message {
      color: white;
      font-size: 10px;
    }
  }

  #input-chat-message {
    margin-left: 2px;
    width: 490px;
  }
}
</style>
