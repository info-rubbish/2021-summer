// import VueRouter from 'vue-router'
import { createRouter, createWebHistory } from 'vue-router'
import home from '@/pages/general/home.vue'
import regist from '@/pages/general/regist.vue'
import login from '@/pages/general/login.vue'
import edit from '@/pages/article/edit.vue'
import read from '@/pages/article/read.vue'
import list from '@/pages/article/list.vue'
import load from '@/pages/article/load.vue'
import find from '@/pages/article/find.vue'
import account from '@/pages/general/account.vue'
var router = createRouter({
    routes: [
        { path: '/login', name: 'login', component: login },
        { path: '/article/find/:query',  component: find },
        { path: '/article/edit/:id',  component: edit },
        { path: '/article/read/:id', component: read},
        { path: '/article/load', name: 'articleload', component: load },
        { path: '/article/list', name: 'articlelist', component: list },
        { path: '/home', name: 'home', component: home },
        { path: '/regist', name: 'regist', component: regist },
        { path: '/account', name: 'account', component: account },
        { path: '/', redirect: '/home' },
    ],
    history: createWebHistory(),
})
export default router
