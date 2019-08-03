import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
    plugins: [createPersistedState()],
    state: {
        // endpoint: '10.8.0.1:9000',
        endpoint: 'localhost:8080',
        dark: false
    },
    getters: {
        dark: (state) => state.dark,
        endpoint: (state) => state.endpoint
    },
    actions: {
        SET_DARK: (context, value) => context.commit('SET_DARK', value),
        SET_ENDPOINT: (context, value) => context.commit('SET_ENDPOINT', value),
    },
    mutations: {
        SET_DARK: (state, value) => state.dark = value,
        SET_ENDPOINT: (state, value) => state.endpoint = value
    },
    strict: process.env.NODE_ENV !== 'production'
})
