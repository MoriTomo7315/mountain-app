import Vuex from 'vuex'
import user from './modules/user'
import posts from './modules/posts'

export default () => new Vuex.Store({
  modules: {
    user,
    posts
  }
})
