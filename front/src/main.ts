import './assets/main.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'

import Vue3Toastify from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { defaultToastOptions } from './utils/toastr';

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(Vue3Toastify, defaultToastOptions);
app.use(router)

app.mount('#app')
