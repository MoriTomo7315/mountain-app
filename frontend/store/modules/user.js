const state = {
  user: null,
  loginStatus: false
}

const getters = {
  user: (state) => state.user
}

const mutations = {
  setUser(state, { user }) {
    state.user = user
  },
  login(state ) {
    state.loginStatus = true
  },
  logout(state) {
    state.user = null,
    state.loginStatus = false
  }
}

const actions = {
  fetchUser({ commit }, user) {
    commit('setUser', {user})
  },
  login({ commit }, user) {
    commit('login', {user})
  },
  logout({ commit }) {
    // firebase.auth().signOut().then(() => {
    //     commit("logout")
    // })
    commit("logout")
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
