import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Vant from 'vant';
import 'vant/lib/index.css';
import api from './api/http'


Vue.use(Vant);
Vue.config.productionTip = false
Vue.prototype.$API = api

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
