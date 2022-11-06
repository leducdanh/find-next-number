import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    drawer: false,
    curUser: null,
  },
  mutations: {
    toggleDrawer: function (state) {
      return state.drawer = !state.drawer;
    },
    setCurUser: function (state, user) {
      state.curUser = user;
    }
  },
  actions: {
    toggleDrawer({ commit }) {
      commit('toggleDrawer');
    },
    setCurUser({ commit }) {
      commit('setCurUser');
    }
  },
  getters: {
    active: (state) => {
      return state.drawer;
    },
    getCurUser: (state) => {
      return state.curUser;
    },
  }
})