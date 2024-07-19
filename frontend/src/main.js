import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

const app = createApp(App)
import 'tdesign-vue-next/es/style/index.css';
app.use(createPinia())
app.use(router)
app.mount('#app')
