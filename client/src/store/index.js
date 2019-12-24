import Vue from 'vue';
import Vuex from 'vuex';
import Network from '../Network';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        status: false,
        email: {}
    },
    getters: {
        getStatus(state) {
            return state.status;
        }
    },
    mutations: {
        setStatus(state) {
            state.status = true;
        },
        setEmail(state, data) {
            state.email = data;
        }
    },
    actions: {
        async pingServer({ commit }) {
            const { status } = await Network.ping();
            if (status === 200) {
                commit('setStatus');
            }
        }
    }
})
