import { createApp } from 'vue'
import router from '@/router'
import App from '@/App.vue'
import axios from 'axios'
import { api } from '@/utils/api.js'
import '@/all.css'
window.axios = axios
window.api = api
// window.axios=await import("axios")
var app = createApp(App);
app.component(App.name, App);
app.use(router).mount('#app');