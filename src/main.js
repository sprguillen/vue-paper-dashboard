import Vue from "vue";
import Vuex from "vuex";
import Toasted from 'vue-toasted';
import App from "./App";
import router from "./router/index";
import store from "./store"

import PaperDashboard from "./plugins/paperDashboard";
import "vue-notifyjs/themes/default.css";

const options = {
	position: 'top-center',
	duration: '2000'
};

Vue.use(Vuex);
Vue.use(Toasted, options);
Vue.use(PaperDashboard);

/* eslint-disable no-new */
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
