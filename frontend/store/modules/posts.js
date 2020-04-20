import axios from 'axios'
// apiのurl
const URL = '/api/posts'

const state = {
  posts: []
}

const getters = {
  posts: (state) => state.posts
}

const mutations = {
  setPosts(state, { posts }) {
    state.posts = posts
  }
}

const actions = {
  async fetchPosts({ commit }) {
    await axios.get(URL).then((responce) => {
      commit('setPosts', {posts: responce.data})
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
