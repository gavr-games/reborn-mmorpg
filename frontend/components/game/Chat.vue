<template>
  <div id="chat" class="rpgui-container framed-golden">
    <div id="chat-messages" class="rpgui-container framed-grey">
      <div class="message" v-for="index in chatMessages.length" :key="'msg'+index">
        {{ chatMessages[index - 1] }}
      </div>
    </div>
    <input id="input-chat-message" placeholder="type your message here" v-model="chatMessage" v-on:keyup.enter="sendChatMessage" autocomplete="off">
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
  },

  beforeDestroy() {
    EventBus.$off("new-chat-message")
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
  }
}
</script>


<style>
#chat {
  width: 500px;
  height: 300px;
  position: absolute;
  bottom: 0;
  left: 0;
  opacity: 0.8;

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
    width: 442px;
  }
}
</style>
