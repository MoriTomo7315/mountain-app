const state = {
  user: null,
  loginStatus: false
}

const getters = {
  user: (state) => state.user,
  isLogin: (state) => state.loginStatus
}

const mutations = {
  setUser(state, { user }) {
    state.user = user
  },
  login(state) {
    state.loginStatus = true
  },
  logout(state) {
    state.loginStatus = false
    state.user = null
  }
}

const actions = {
  fetchUser({ commit }, user) {
    commit('setUser', {user})
  },
  login({ commit }) {
    commit('login')
  },
  logout({ commit }) {
    firebase.auth().signOut().then(() => {
        commit("logout")
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
