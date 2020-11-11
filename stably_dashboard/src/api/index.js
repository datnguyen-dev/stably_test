import axios from 'axios'

axios.interceptors.request.use((cfg) => {
  return cfg
}, (error) => {
  return Promise.reject(error)
})

axios.interceptors.response.use((response) => {
  return response
}, (error) => {
  // Do something with response error
  return Promise.reject(error)
})

const baseUrl = (process.env.VUE_APP_API_URL || '')

export default {
  baseUrl () {
    return baseUrl
  },
  request (method, uri, data = null) {
    if (!method) {
      // eslint-disable-next-line
      console.error('API function call requires method argument')
      return
    }
    if (!uri) {
      // eslint-disable-next-line
      console.error('API function call requires uri argument')
      return
    }
    let url = this.baseUrl() + uri
    return axios({ method, url, data })
  }
}
