import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import store from './store';
import GAuth from 'vue-google-oauth2'

Vue.config.productionTip = false
const gauthOption = {
  clientId: '137722493716-ndhcrcrs7i33fni3nspmchtmsji0rb2d.apps.googleusercontent.com',
  scope: 'profile email',
  prompt: 'consent',
  plugin_name: 'findnextnumber',
}
Vue.use(GAuth, gauthOption)
new Vue({
  vuetify,
  store,
  render: h => h(App)
}).$mount('#app')
