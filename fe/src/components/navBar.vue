<template>
    <div
        class="lg:sticky lg:h-auto lg:w-full h-full fixed top-0 rounded-2xl"
        id="app1"
    >
        <div class="absolute z-50 lg:hidden">
            <!--lg+ hide-->
            <button @click="show = !show">
                <MenuIcon class="w-10 h-10" />
            </button>
        </div>
        <ul
            class="
                flex
                lg:w-full lg:h-auto lg:transform-none lg:p-2 lg:flex-row
                flex-col
                p-10
                h-full
                justify-around
                text-2xl
                transition-transform
                transform
                bg-gray-300
                w-60
            "
            :class="show ? 'translate-x-0' : '-translate-x-full'"
        >
            <input
                class="focus:outline-none rounded-xl w=4/12 p-2 block"
                placeholder="尋找課程"
            />
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/home" class="whitespace-nowrap"
                    >&nbsp;主頁&nbsp;</router-link
                >
            </li>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/login" class="whitespace-nowrap"
                    >&nbsp;登入&nbsp;</router-link
                >
            </li>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/login" class="whitespace-nowrap"
                    >&nbsp;登出&nbsp;</router-link
                >
            </li>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/regist" class="whitespace-nowrap"
                    >&nbsp;註冊&nbsp;</router-link
                >
            </li>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/article/list" class="whitespace-nowrap"
                    >&nbsp;課程&nbsp;</router-link
                >
            </li>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="#" class="whitespace-nowrap">
                    &nbsp;帳號&nbsp;
                </router-link>
            </li>
        </ul>
    </div>
</template>
<script>
import 'vue'
import { MenuIcon } from '@heroicons/vue/solid'
import { api } from '@/utils/api.js'
import localStorage from '@/utils/localstorage.js'

export default {
    name: 'navBar',
    methods: {
        async logout(){
            var res = await new api().Logout()
            if(res) this.$router.push('/home')
        }
    },
    data() {
        var api_ = new api()
        if (api_.CheckTTL())
        api_.GetSelfInfo().then(console.log)
        else console('ttl!')
        return { show: false }
    },
    components: {
        MenuIcon,
    },
}
// function app1
// w-full
</script>
<style>
ul.list > li {
    display: inline-block;
}
</style>
