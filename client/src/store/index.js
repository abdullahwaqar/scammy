import Vue from 'vue';
import Vuex from 'vuex';
import Network from '../Network';
import router from '../router';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        status: false,
        email: {}
    },
    getters: {
        getStatus(state) {
            return state.status;
        },
        getEmail(state) {
            return state.email;
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
        },
        async getEmailDetails({ commit }, email) {
            const response = await Network.getEmailDEtails(email);
            if (response.error) {
                return;
            }
            commit('setEmail', response);
            router.push('/scan');
            console.log(response)
        }
    }
})
