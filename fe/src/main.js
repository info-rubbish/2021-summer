import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import 'axios'
import './all.css'
window.axios=await import("axios")
var app=createApp(App);
app.component(App.name,App);
app.use(router).mount('#app');
