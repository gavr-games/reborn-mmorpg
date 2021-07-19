export const state = () => ({
  selectedCharacterId: null
})

export const mutations = {
  set(state, id) {
    state.selectedCharacterId = id
  }
}
