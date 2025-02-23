<script setup lang="ts">
import { h, ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from './stores/user'
import {
  NConfigProvider,
  NLayout,
  NLayoutHeader,
  NLayoutContent,
  NMenu,
  NButton,
  NGradientText,
  NLoadingBarProvider,
  NMessageProvider,
  NNotificationProvider,
  NDialogProvider,
  darkTheme,
  useOsTheme
} from 'naive-ui'
import {
  HomeOutline as HomeIcon,
  DocumentTextOutline as ArticleIcon,
  HelpCircleOutline as QuestionIcon
} from '@vicons/ionicons5'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

// 主题设置
const theme = computed(() => {
  const osTheme = useOsTheme()
  return osTheme.value === 'dark' ? darkTheme : null
})

// 在组件挂载时初始化用户状态
onMounted(async () => {
  await userStore.initToken()

  // 只有在非登录相关页面且未登录时才重定向
  const isLoginPage = location.pathname.startsWith('/login')
  const isCallback = location.pathname === '/login/callback'
  if (!userStore.token && !isLoginPage && !isCallback) {
    router.push('/login')
  }
})

// 导航菜单
const activeKey = ref(route.path)

// 监听路由变化更新激活的菜单项
watch(
  () => route.path,
  (newPath) => {
    activeKey.value = newPath
  }
)

const menuOptions = [
  {
    label: '首页',
    key: '/',
    icon: renderIcon(HomeIcon)
  },
  {
    label: '文章',
    key: '/articles',
    icon: renderIcon(ArticleIcon)
  },
  {
    label: '问题',
    key: '/questions',
    icon: renderIcon(QuestionIcon)
  }
]

function renderIcon(icon: any) {
  return () => h('div', { style: { marginRight: '6px' } }, h(icon))
}

// 处理菜单点击
const handleMenuClick = (key: string) => {
  router.push(key)
}

// 登出处理
const handleLogout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<template>
  <n-config-provider :theme="theme">
    <n-loading-bar-provider>
      <n-message-provider>
        <n-notification-provider>
          <n-dialog-provider>
            <n-layout position="absolute">
              <n-layout-header v-if="userStore.token" bordered>
                <div class="nav-header">
                  <div class="nav-brand">
                    <n-gradient-text type="success">笔记同步系统</n-gradient-text>
                  </div>
                  <div class="nav-menu">
                    <n-menu
                      v-model:value="activeKey"
                      mode="horizontal"
                      :options="menuOptions"
                      @update:value="handleMenuClick"
                    />
                  </div>
                  <div class="nav-actions">
                    <n-button
                      @click="handleLogout"
                      quaternary
                      type="error"
                    >
                      退出登录
                    </n-button>
                  </div>
                </div>
              </n-layout-header>
              <n-layout-content content-style="padding: 24px;">
                <router-view></router-view>
              </n-layout-content>
            </n-layout>
          </n-dialog-provider>
        </n-notification-provider>
      </n-message-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: v-sans, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
    sans-serif;
}

.nav-header {
  padding: 8px 24px;
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  gap: 24px;
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
}

.nav-menu {
  justify-self: center;
  min-width: 400px;
}

.nav-actions {
  display: flex;
  gap: 12px;
  align-items: center;
}

.n-layout-content {
  min-height: calc(100vh - 64px);
}
</style>
