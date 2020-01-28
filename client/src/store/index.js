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
        },
        getEmailHeader(state) {
            return state.email.email_header.split(',').join('\n');
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
        },
        async incrementReportCount({ commit }) {
            const response = await Network.incrementEmailReport(this.state.email.id);
            commit('setEmail', response);
        },
        async reportNewScam({ commit, dispatch }, { email, emailHeader }) {
            //* Check is the email already exists
            const emailCheck = await Network.getEmailDEtails(email);
            if (!emailCheck.error) {
                //* Email already exists, just increment it
                const response = await Network.incrementEmailReport(emailCheck.id);
                commit('setEmail', response);
            } else {
                const response = await Network.postNewScamEmail(email, emailHeader);
                commit('setEmail', response);
            }
        }
    }
})
