import GameController from "~/plugins/game/game_controller"

export default ({ app }, inject) => {
  inject('game', GameController)
}
