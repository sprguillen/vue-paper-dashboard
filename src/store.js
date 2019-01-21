import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		authenticated: false,
		currentDashboardTime: ''
	},
	mutations: {
		login(state) {
			state.authenticated = true;
		},
		logout(state) {
			state.authenticated = false;
		},
		setCurrentTime(state, currentTime) {
			state.currentDashboardTime = currentTime;
		}
	},
	actions: {
		loginAction({commit, state}, email) {
			localStorage.email = email;
			commit('login');
		},
		logoutAction({commit, state}) {
			localStorage.clear();
			commit('logout');
		}
	}
});
