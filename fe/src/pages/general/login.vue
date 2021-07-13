<template>
    <div>
        <navbar />
        <form v-on:submit.prevent="login">
            <div class="grid grid-cols-2 border-blue">
                <label class="p-1 border-t-4 border-fuchsia-600 leading-10">帳戶</label><input type="text" class="focus:outline-none p-1 border-t-4 border-fuchsia-600 leading-10" placeholder="輸入帳號"> 
                <label class="p-1 border-t-4 border-b-4 border-fuchsia-600 leading-10">密碼</label><input type="password" class="focus:outline-none p-1 leading-10 border-t-4 border-b-4 border-fuchsia-600" placeholder="輸入密碼"> 
            </div>
            <div><input type="submit" class="hover:bg-green-300 mt-4"></div>
        </form>
        <!--<router-link to="/regist"><span>註冊</span></router-link>-->
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
import { api } from '@/utils/api.js'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async login(e) {
            const name = e.target[0].value
            const password = e.target[1].value
            var res = await new api().Login({ name, password })
            if(res) this.$router.push('/home')
        },
    }, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
</style>
