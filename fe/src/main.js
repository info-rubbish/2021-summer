import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import './all.css'
var app=createApp(App);
app.component(App.name,App);
app.use(router).mount('#app');