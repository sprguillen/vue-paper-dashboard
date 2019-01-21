import Vue from "vue";
import VueRouter from "vue-router";
import routes from "./routes";
import store from "../store";
Vue.use(VueRouter);

// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: "active"
});

const email = localStorage.email;

router.beforeEach((to, from, next) => {
	if (to.name === 'dashboard') {
		if (email) {
			store.commit('login');
			next();
		} else {
			next('/login');
		}
	} else if (to.name === 'login') {
		if (email) {
			next('/dashboard');
		} else {
			next();
		}
	}
});

export default router;
