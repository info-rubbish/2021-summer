<template>
    <div>
        <navbar />
        <span
                @click="logout"
                class="
                    place-self-center
                    cursor-pointer
                    w-24
                    leading-relaxed
                    text-white text-xl
                    font-extrabold
                    rounded-md
                    flex
                    items-center
                    justify-center
                    bg-blue-500
                    mx-2
                    my-6
                    p-2
                    float-right
                "
                >登出</span
            >
                        <router-link
                to="/article/load"
                v-show="user.permission == 2"
                class="
                    place-self-center
                    w-24
                    leading-relaxed
                    text-white text-xl
                    font-extrabold
                    rounded-md
                    flex
                    items-center
                    justify-center
                    bg-blue-500
                    mx-2
                    my-6
                    p-2
                    float-right
                "
                >管理文章</router-link
            >
        <div class="clear"></div>
        <div class="flex mx-6 my-10 flex-wrap justify-around ">

            <form class="bg-gray-200 w-72 m-4" @submit.prevent="deleteAccount">
                <h1 class="text-2xl text-center p-2">刪除帳號</h1>
                <div class="grid grid-cols-2">
                    <label for="config-user-name-delete">輸入名稱來確認</label>
                    <input type="text" id="config-user-name-delete" />
                </div>
                <button type="submit" class="bg-red-500 p-2 my-4">
                    確認刪除
                </button>
            </form>
            <form class="bg-gray-200 w-72 m-4" @submit.prevent="editAccount">
                <h1 class="text-2xl text-center p-2">修改使用者資訊</h1>
                <div class="grid grid-cols-2">
                    <label for="config-user-name">名稱</label
                    ><input type="text" id="config-user-name" />
                    <label for="config-user-password">密碼</label
                    ><input type="password" id="config-user-password" />
                    <label for="config-user-password-check">確認密碼</label
                    ><input type="password" id="config-user-password-check" />
                </div>
                <button type="submit" class="bg-blue-500 p-2 my-4">
                    確認修改
                </button>
            </form>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
import { mapState } from 'vuex'
export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async logout() {
            await this.$store.dispatch('Logout')
            this.$router.push('/home')
        },
    },
    async mounted() {},
    data() {
        return {}
    },
    computed: { ...mapState(['user']) },
}
</script>
<style>
.clear{
    clear: both;
}</style>
