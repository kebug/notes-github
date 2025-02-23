import axios from 'axios'
import { useUserStore } from '../stores/user'

const request = axios.create({
  baseURL: 'http://localhost:8080/api'
})

// 请求拦截器添加token
request.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

export default request 