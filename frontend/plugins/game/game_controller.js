import SurfaceController from './objects/surface/surface_controller'
import RockController from './objects/rock/rock_controller'
import TreeController from './objects/tree/tree_controller'
import PlantController from './objects/plant/plant_controller'
import ItemController from './objects/item/item_controller'
import CharacterController from './objects/character/character_controller'
import MobController from './objects/mob/mob_controller'
import NpcController from './objects/npc/npc_controller'
import ClaimAreaController from './objects/claim_area/claim_area_controller'
import SelectCoordsController from './select_coords/select_coords_controller'
import GameConnnection from './game_connection'
import { EventBus } from '~/plugins/game/event_bus'
import ChatController from '~/plugins/game/chat/chat_controller'
import GameObserver from '~/plugins/game/game_observer'

const PING_INTERVAL = 20000

class GameController {
  constructor () {
    this.gameObjects = []
    this.token = null
    this.characterId = null
    this.targetObjId = null
    this.ping = [0, 0, 0]
    this.pingIndex = 0
    this.controls = {
      w: false,
      a: false,
      s: false,
      d: false,
      lastCmd: ''
    }
    this.initGameObjectsHandler = (gameObjects) => {
      this.initGameObjects(gameObjects)
      this.startPing()
    }
    this.keyUpHandler = (key) => {
      this.handleKeyUp(key)
    }
    this.keyDownHandler = (key) => {
      this.handleKeyDown(key)
    }
    this.addObjectHandler = (gameObj) => {
      this.createObject(gameObj)
    }
    this.updateObjectHandler = (gameObj) => {
      if (this.gameObjects[gameObj.Id]) {
        this.gameObjects[gameObj.Id].update(gameObj)
      } else {
        this.createObject(gameObj)
      }
      if (gameObj.Properties && gameObj.Properties.container_id) {
        EventBus.$emit('update-item-in-container', {
          container_id: gameObj.Properties.container_id,
          item: gameObj
        })
      }
    }
    this.removeObjectHandler = (gameObj) => {
      if (this.gameObjects[gameObj.Id]) {
        this.gameObjects[gameObj.Id].remove()
        this.gameObjects[gameObj.Id] = null
      }
    }
    this.selectTargetHandler = (targetObj) => {
      this.targetObjId = targetObj.id
      if (this.gameObjects[targetObj.id]) {
        this.gameObjects[targetObj.id].selectAsTarget()
      }
    }
    this.deselectTargetHandler = (targetObj) => {
      this.targetObjId = null
      if (this.gameObjects[targetObj.id]) {
        this.gameObjects[targetObj.id].deselectAsTarget()
      }
    }
    this.addObjectsHandler = (data) => {
      data.objects.forEach((gameObj) => {
        EventBus.$emit('add_object', gameObj)
      })
    }
    this.removeObjectsHandler = (data) => {
      data.objects.forEach((gameObj) => {
        EventBus.$emit('remove_object', gameObj)
      })
    }
    this.setPingHandler = (data) => {
      this.ping[this.pingIndex] = data.server_time - data.start_time
      this.pingIndex++
      if (this.pingIndex === 3) {
        this.pingIndex = 0
      }
    }
    this.performGameAction = (action) => {
      GameConnnection.sendCmd(action.cmd, action.params)
    }
    EventBus.$on('init_game', this.initGameObjectsHandler)
    EventBus.$on('keyup', this.keyUpHandler)
    EventBus.$on('keydown', this.keyDownHandler)
    EventBus.$on('add_object', this.addObjectHandler)
    EventBus.$on('update_object', this.updateObjectHandler)
    EventBus.$on('remove_object', this.removeObjectHandler)
    EventBus.$on('perform-game-action', this.performGameAction)
    EventBus.$on('select_target', this.selectTargetHandler)
    EventBus.$on('deselect_target', this.deselectTargetHandler)
    EventBus.$on('add_objects', this.addObjectsHandler)
    EventBus.$on('remove_objects', this.removeObjectsHandler)
    EventBus.$on('ping_info', this.setPingHandler)
  }

  init (token, characterId) {
    this.token = token
    this.characterId = characterId

    GameConnnection.init(token, characterId)
    GameObserver.init()
    ChatController.init(token, characterId)
    new SelectCoordsController() // eslint-disable-line no-new
  }

  startPing () {
    GameConnnection.sendCmd('get_ping', { start_time: Date.now() })
    setInterval(() => {
      GameConnnection.sendCmd('get_ping', { start_time: Date.now() })
    }, PING_INTERVAL)
  }

  initGameObjects (gameObjects) {
    gameObjects.visible_objects.forEach((gameObj) => {
      this.createObject(gameObj)
    })
  }

  createObject (gameObj) {
    if (this.gameObjects[gameObj.Id] != null) {
      return
    }
    // don't show not visible objects
    if (gameObj.Properties.visible == null || !gameObj.Properties.visible) {
      return
    }
    switch (gameObj.Type) {
      case 'surface':
        this.gameObjects[gameObj.Id] = new SurfaceController(gameObj)
        break
      case 'player':
        this.gameObjects[gameObj.Id] = new CharacterController(gameObj, this.characterId)
        break
      case 'rock':
        this.gameObjects[gameObj.Id] = new RockController(gameObj)
        break
      case 'tree':
        this.gameObjects[gameObj.Id] = new TreeController(gameObj)
        break
      case 'plant':
        this.gameObjects[gameObj.Id] = new PlantController(gameObj)
        break
      case 'mob':
        this.gameObjects[gameObj.Id] = new MobController(gameObj)
        break
      case 'npc':
        this.gameObjects[gameObj.Id] = new NpcController(gameObj)
        break
      case 'claim':
        if (gameObj.Properties.kind === 'claim_area') {
          this.gameObjects[gameObj.Id] = new ClaimAreaController(gameObj)
        } else {
          this.gameObjects[gameObj.Id] = new ItemController(gameObj)
        }
        break
      default:
        this.gameObjects[gameObj.Id] = new ItemController(gameObj)
        break
    }

    if (this.targetObjId === gameObj.Id) {
      this.selectTargetHandler(gameObj.Properties)
    }
  }

  handleKeyUp (key) {
    if (['w', 'a', 's', 'd'].includes(key)) {
      this.controls[key] = false
      this.moveCharacter()
    }
  }

  handleKeyDown (key) {
    if (['w', 'a', 's', 'd'].includes(key)) {
      this.controls[key] = true
      this.moveCharacter()
    }
    if (key === '1') { // Hit the target
      GameConnnection.sendCmd('melee_hit')
    }
  }

  moveCharacter () {
    let directionSum = 0
    if (this.controls.w) { directionSum += 1 }
    if (this.controls.a) { directionSum += 10 }
    if (this.controls.s) { directionSum += 100 }
    if (this.controls.d) { directionSum += 1000 }
    let cmd = 'stop'
    switch (directionSum) {
      case 1:
        cmd = 'move_north_west'
        break
      case 10:
        cmd = 'move_south_west'
        break
      case 100:
        cmd = 'move_south_east'
        break
      case 1000:
        cmd = 'move_north_east'
        break
      case 11:
        cmd = 'move_west'
        break
      case 110:
        cmd = 'move_south'
        break
      case 1100:
        cmd = 'move_east'
        break
      case 1001:
        cmd = 'move_north'
        break
    }
    if (cmd !== this.controls.lastCmd) {
      // EventBus.$emit('emulate-character-move', {
      //   cmd,
      //   ping: this.ping.reduce((a, b) => a + b) / this.ping.length
      // }) // this helps to predict character movement before server answer and get rid of "jumps" (ping)
      GameConnnection.sendCmd(cmd)
      this.controls.lastCmd = cmd
    }
  }

  destroy () {
    ChatController.destroy()
    EventBus.$off('init_game', this.initGameObjectsHandler)
    EventBus.$off('keyup', this.keyUpHandler)
    EventBus.$off('keydown', this.keyDownHandler)
    EventBus.$off('add_object', this.addObjectHandler)
    EventBus.$off('update_object', this.updateObjectHandler)
    EventBus.$off('remove_object', this.removeObjectHandler)
    EventBus.$off('perform-game-action', this.performGameAction)
    EventBus.$off('select_target', this.selectTargetHandler)
    EventBus.$off('deselect_target', this.deselectTargetHandler)
    EventBus.$off('add_objects', this.addObjectsHandler)
    EventBus.$off('remove_objects', this.removeObjectsHandler)
    EventBus.$off('ping_info', this.setPingHandler)
  }
}

const gameController = new GameController()

export default gameController
