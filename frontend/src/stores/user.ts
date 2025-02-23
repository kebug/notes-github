import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'
import type { User } from '../types'

export const useUserStore = defineStore('user', () => {
  const token = ref<string | null>(null)
  const user = ref<User | null>(null)

  const setToken = async (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
    await loadUser() // 设置token后自动加载用户信息
  }

  const setUser = (newUser: User) => {
    user.value = newUser
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    // 清除 GitHub OAuth 授权
    window.location.href = 'https://github.com/settings/applications'
  }

  // 从后端加载用户信息
  const loadUser = async () => {
    if (!token.value) return

    try {
      const response = await axios.get('http://localhost:8080/api/user', {
        headers: {
          Authorization: `Bearer ${token.value}`
        }
      })
      user.value = response.data
    } catch (error) {
      console.error('Failed to load user:', error)
      logout() // 如果获取用户信息失败，清除登录状态
    }
  }

  // 初始化时从localStorage加载token并获取用户信息
  const initToken = async () => {
    const savedToken = localStorage.getItem('token')
    if (savedToken) {
      token.value = savedToken
      try {
        await loadUser() // 加载用户信息
      } catch (error) {
        console.error('Failed to load user:', error)
        // 如果加载用户信息失败，清除token
        logout()
      }
    }
  }

  return {
    token,
    user,
    setToken,
    setUser,
    logout,
    initToken,
    loadUser
  }
}) 