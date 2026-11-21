import App from './App'
import store from './store';

// #ifndef VUE3
import Vue from 'vue';
import uView from '@/uni_modules/uview-ui'
Vue.use(uView)
import storea from '@/common/store'
Vue.prototype.$storea = storea;
import {
	stop,
	move
} from '@/common/utils.js'
Vue.prototype.stop = stop
Vue.prototype.move = move
Vue.config.productionTip = false
App.mpType = 'app'

const app = new Vue({
    ...App,
	store
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
export function createApp() {
  const app = createSSRApp(App)
  return {
    app
  }
}
// #endif