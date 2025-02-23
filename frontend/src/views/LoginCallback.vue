<template>
  <div class="login-callback">
    <h2>正在登录...</h2>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '../stores/user'
import { useMessage } from '../utils/message'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const message = useMessage()

onMounted(async () => {
  const code = route.query.code as string
  try {
    await userStore.handleCallback(code)
    message.success('登录成功')
    router.push('/')
  } catch (error) {
    message.error('登录失败：' + error)
    router.push('/login')
  }
})
</script>

<style scoped>
.login-callback {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style> 