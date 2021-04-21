import Vue from 'vue'
import Qs from 'qs'
import {
  Message
} from 'element-ui'
// Vue.use(Message);
Vue.component(Message.name, Message)

import axios from 'axios'
import {
  getLocal
} from '../common/utils'
export const baseURL = '/api'
// export const baseURL = "../";
axios.defaults.timeout = 8000
axios.defaults.baseURL = process.env.API_HOST
axios.defaults.headers.post['Content-Type'] = 'application/json'
axios.defaults.headers.common['locale'] = getLocal()

axios.interceptors.request.use(
  (config) => {
    // if (config.method === 'post') {
    //   // config.data = JSON.stringify({
    //   //   ...config.data,
    //   // });
    // } else if (config.method === 'get') {}
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

axios.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response && error.response.status) {
      Message.error(error.response.status + ' ' + error.response.statusText)
    } else {
      Message.error(error.message)
    }
    return Promise.reject(error)
  }
)

export function get(url, params = {}) {
  return new Promise((resolve, reject) => {
    axios
      .get(url, {
        params: params
      })
      .then((response) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, data)
      .then((response) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}
export function postForm(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios({
      method: 'POST',
      headers: { 'content-type': 'application/x-www-form-urlencoded' },
      data: Qs.stringify(data),
      url
    })
      .then((response) => {
        resolve(response.data)
      })
      .catch((error) => {
        reject(error)
      })
  })
}
