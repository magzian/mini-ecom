import './assets/main.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { store } from '../store'

const app = createApp(App)
app.config.globalProperties.$store = store
app.use(router)
app.mount('#app')
