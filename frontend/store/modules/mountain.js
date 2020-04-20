import axios from 'axios'
// apiのurl
const URL = '/api/mountains'

const state = {
    mountains: []
}

const getters = {
    mountains: (state) => state.mountains
}

const mutations = {
    setMountains (state, { mountains }) {
      state.mountains = mountains
    }
}

const actions = {
    async fetchMountains({ commit }) {
      await axios.get(URL).then((responce) => {
        commit('setMountains', {mountains: responce.data})
      })
    }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
