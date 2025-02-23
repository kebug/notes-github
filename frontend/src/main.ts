import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router'
import { setupNaive } from './components'

// 导入naive-ui样式
import 'vfonts/Lato.css'
import 'vfonts/FiraCode.css'

const app = createApp(App)

setupNaive(app)
app.use(createPinia())
app.use(router)
app.mount('#app')
