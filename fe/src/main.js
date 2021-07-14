import { createApp } from 'vue'
import router from '@/router'
import App from '@/App.vue'
import axios from 'axios'
import '@/all.css'
import {createStore} from 'vuex'
import api from '@/utils/api.js'
window.axios = axios 
// window.axios=await import("axios")

var app = createApp(App);
var store=createStore(api);

app.component(App.name, App);
app.use(router).use(store).mount('#app');
