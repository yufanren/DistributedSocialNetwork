import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'

// Buefy
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
Vue.use(Buefy)

// mavon-editor
import MavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(MavonEditor)

// global css
import './assets/index.css'

Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
