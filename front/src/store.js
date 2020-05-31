import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    msg: "メッセージ"
  },
  mutations: {
    updateMessage (state, message) {
      state.msg = message
    }
  },
  actions: {
    getMessage ({ commit }) {
      axios.get('http://localhost:3000/user/get/1')
      .then(response => {
        if (response.status === 200) {
          commit("updateMessage", response.data);
        }
      })
    }
  }
})

// ストアをエクスポート
export default store
