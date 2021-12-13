import Vue from 'vue';
import Element from 'element-ui'
import router from '../router'
import axios from "axios";
import store from "../store"

let config = {
  baseURL: "http://localhost:8012"
  // timeout: 60 * 1000, // Timeout
  // withCredentials: true, // Check cross-site Access-Control
}

const _axios = axios.create(config)

_axios.interceptors.request.use(config => {
  if(store.getters.token) {
    config.headers['Authorization'] = store.getters.token
  }
  return config
})

_axios.interceptors.response.use(response => {
    const res = response.data
    console.log('=================')
    console.log(res)
    console.log('=================')
    if(res.code == 200) {
      return response
    } else {
      Element.Message({
        message: response.data.message,
        type: 'error',
        duration: 2 * 1000
      })
      return Promise.reject(response.data.message)
    }
  },
  error => {
    if(error.response.data) {
      error.message = error.response.data.msg
    }
    if(error.response.status == 401) {
      store.commit("REMOVE_INFO")
      router.push({
        path: '/login'
      })
      error.message = "Please login again"
    }
    else if(error.response.status == 403) {
      error.message = "No permission to access"
    }
    Element.Message({
      message: error.message,
      type: 'error',
      duration: 3 * 1000
    })
    return Promise.reject(error)
  }
)

Plugin.install = function(Vue, options) {
  Vue.axios = _axios;
  window.axios = _axios;
  Object.defineProperties(Vue.prototype, {
    axios: {
      get() {
        return _axios;
      }
    },
    $axios: {
      get() {
        return _axios;
      }
    }
  })
}

Vue.use(Plugin)

export default Plugin
