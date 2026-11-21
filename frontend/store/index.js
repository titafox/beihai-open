import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);

import tabbar from './modules/tabbar';
import user from './modules/user';
import pay from './modules/pay';

const store = new Vuex.Store({
  modules: {
    tabbar,
		user,
		pay
  }
});

export default store;