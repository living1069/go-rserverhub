import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store/'
import Vuelidate from 'vuelidate'
import events from '@/plugins/events'
import lodash from '@/plugins/lodash'
import api from '@/plugins/api'
import RequestLoader from "@/components/shared/RequestLoader";

require('@/styles/helpers.css')

Vue.component('request-loader', RequestLoader);
Vue.config.productionTip = false
Vue.use(Vuelidate)
Vue.use(events)
Vue.use(lodash)
Vue.use(api)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
