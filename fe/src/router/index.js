import { createRouter, createWebHistory } from 'vue-router'
import home from '@/pages/general/home.vue'
import regist from '@/pages/general/regist.vue'
import login from '@/pages/general/login.vue'
import edit from '@/pages/article/edit.vue'
import read from '@/pages/article/read.vue'
import list from '@/pages/article/list.vue'
import load from '@/pages/article/load.vue'
var router = createRouter({
    routes: [
        { path: '/login', name: 'login', component: login },
        { path: '/article/edit/:id', name: 'articleedit', component: edit },
        { path: '/article/read/:id', name: 'articleedit', component: read},
        { path: '/article/load', name: 'articleload', component: load },
        { path: '/article/list', name: 'articlelist', component: list },
        { path: '/home', name: 'home', component: home },
        { path: '/regist', name: 'regist', component: regist },
        { path: '/', redirect: '/home' },
    ],
    history: createWebHistory(),
})
export default router
