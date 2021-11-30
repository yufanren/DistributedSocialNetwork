import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: localStorage.getItem("token"),
    user: JSON.parse(localStorage.getItem("user"))
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
      localStorage.setItem("token", token)
    },
    SET_USER: (state, user) => {
      state.user = user
      localStorage.setItem("user", JSON.stringify(user))
    },
    REMOVE_INFO: (state) => {
      state.token = ''
      state.user = ''
      localStorage.removeItem("token")
      localStorage.removeItem("user")
    }
  },
  getters: {
    token: state => {
      return state.token
    },
    user: state => {
      return state.user
    }
  },
  actions: {
  },
  modules: {
  }
})
