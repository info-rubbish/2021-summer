<template>
    <div>
        <navbar />
        <div class="flex justify-center m-12">
            <span @click="deleteAccount" class="cursor-pointer w-30 leading-relaxed text-white text-2xl font-extrabold rounded-md flex items-center justify-center bg-indigo-500 m-2 p-8">刪除帳號</span>
            <router-link to="/article/load" class="w-30 leading-relaxed text-white text-2xl font-extrabold rounded-md flex items-center justify-center bg-indigo-500 m-2 p-8">管理文章</router-link>
            <span @click="logout" class="cursor-pointer w-30 leading-relaxed text-white text-2xl font-extrabold rounded-md flex items-center justify-center bg-indigo-500 m-2 p-8">登出</span>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async deleteAccount() {
            if(confirm('確認刪除帳號？您的資料將不會被保留')==true){
                await this.$store.dispatch('DeleteSelf')
                this.$router.push('/login')
            }
        },
        async logout() {
            await this.$store.dispatch('Logout')
                this.$router.push('/home')
        }
    },
    async mounted() {
        if (!(await this.$store.dispatch('CheckTTL')))
            this.$router.push('/login')
    },
    computed: {},
}
</script>

<style></style>
