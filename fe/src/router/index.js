import { createRouter, createWebHistory } from 'vue-router'
import Home from '/src/pages/Home.vue'
import Regist from '/src/pages/Regist.vue'
import Login from '/src/pages/Login.vue'
var router = createRouter({
    routes: [
        { path: '/login', name: 'Login', component: Login },
        { path: '/home', name: 'Home', component: Home },
        { path: '/regist', name: 'regist', component: Regist },
        { path: '/', redirect: '/home' },
    ],
    history: createWebHistory(),
})
export default router
