import { EventBus } from '~/plugins/game/event_bus'

class ChatController {
  constructor () {
    this.conn = null
    this.sendChatMessage = (message) => {
      if (this.conn) {
        this.conn.send(message)
      }
    }
    EventBus.$on('send-chat-message', this.sendChatMessage)
  }

  init (token, characterId) {
    if (window.WebSocket) {
      this.conn = new WebSocket('ws://' + document.location.host + '/chat/ws?token=' + token + '&character_id=' + characterId)
      /* this.conn.onclose = function (evt) {
        console.log('Chat ws connection is closed')
      } */
      this.conn.onmessage = function (evt) {
        const messages = evt.data.split('\n')
        for (let i = 0; i < messages.length; i++) {
          EventBus.$emit('new-chat-message', messages[i])
        }
      }
    }
  }

  destroy () {
    EventBus.$off('send-chat-message')
  }
}

const chatController = new ChatController()

export default chatController
