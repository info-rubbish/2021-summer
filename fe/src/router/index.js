import { createRouter, createWebHistory } from 'vue-router'
import Home from '/src/pages/Home.vue'
var router = createRouter({
    routes: [{ path: '/home',name:'home', component: Home }],
    history: createWebHistory(),
})
export default router