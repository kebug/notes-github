import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue';
import Articles from '../views/Articles.vue';
import Editor from '../views/Editor.vue';
import Questions from '../views/Questions.vue';
import QuestionEditor from '../views/QuestionEditor.vue';
import Login from '../views/Login.vue';
import LoginCallback from '../views/LoginCallback.vue';
import { useUserStore } from '../stores/user'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/articles',
      component: Articles
    },
    {
      path: '/editor',
      component: Editor
    },
    {
      path: '/questions',
      component: Questions
    },
    {
      path: '/question-editor',
      component: QuestionEditor
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '/login/callback',
      component: LoginCallback
    }
  ]
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  const isLoginPage = to.path.startsWith('/login')
  
  if (!userStore.token && !isLoginPage) {
    // 未登录且不是登录页面，重定向到登录页
    next('/login')
  } else if (userStore.token && isLoginPage) {
    // 已登录但访问登录页，重定向到首页
    next('/')
  } else {
    // 其他情况正常放行
    next()
  }
})

export default router 