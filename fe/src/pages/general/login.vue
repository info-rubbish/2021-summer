<template>
    <navbar />
    <div class="flex justify-center items-center p-10">
        <div class="block w-72 bg-gray-200">
            <div class="flex justify-start p-2">
                <button
                    class="hover:bg-blue-300 bg-gray-300 p-2 my-2"
                    @click="tab = 0"
                    :class="tab == 0 ? 'bg-gray-400' : ''"
                >
                    登入
                </button>
                <button
                    class="hover:bg-blue-300 bg-gray-300 p-2 my-2"
                    @click="tab = 1"
                    :class="tab == 1 ? 'bg-gray-400' : ''"
                >
                    註冊
                </button>
            </div>
            <form v-on:submit.prevent="login" v-show="tab == 0">
                <div class="grid grid-cols-2 border-blue">
                    <label class="p-1 border-t-4 border-fuchsia-600 leading-10">
                        帳戶</label
                    ><input
                        type="text"
                        class="
                            focus:outline-none
                            p-1
                            border-t-4 border-fuchsia-600
                            leading-10
                        "
                        placeholder="輸入帳號"
                    />
                    <label
                        class="
                            p-1
                            border-t-4 border-b-4 border-fuchsia-600
                            leading-10
                        "
                        >密碼</label
                    ><input
                        type="password"
                        class="
                            focus:outline-none
                            p-1
                            leading-10
                            border-t-4 border-b-4 border-fuchsia-600
                        "
                        placeholder="輸入密碼"
                    />
                </div>
                <div class="flex justify-start">
                    <button
                        type="submit"
                        class="hover:bg-blue-300 bg-gray-300 p-2 mt-4"
                    >
                       登入
                    </button>
                </div>
            </form>
            <form v-on:submit.prevent="regist" v-show="tab == 1">
                <div class="grid grid-cols-2 border-blue">
                    <label class="p-1 border-t-4 border-fuchsia-600 leading-10">
                        帳戶</label
                    ><input
                        type="text"
                        class="
                            focus:outline-none
                            p-1
                            border-t-4 border-fuchsia-600
                            leading-10
                        "
                        placeholder="輸入帳號名稱"
                    />
                    <label class="p-1 border-t-4 border-fuchsia-600 leading-10"
                        >密碼</label
                    ><input
                        type="password"
                        class="
                            focus:outline-none
                            p-1
                            leading-10
                            border-t-4 border-fuchsia-600
                        "
                        placeholder="輸入密碼"
                    />
                    <label
                        class="
                            p-1
                            border-t-4 border-b-4 border-fuchsia-600
                            leading-10
                        "
                        >再輸入一次</label
                    >
                    <input
                        type="password"
                        class="
                            focus:outline-none
                            p-1
                            leading-10
                            border-t-4 border-b-4 border-fuchsia-600
                        "
                        placeholder="再輸入一次"
                    />
                </div>
                <div class="flex justify-start">
                    <button
                        type="submit"
                        class="hover:bg-blue-300 bg-gray-300 p-2 mt-4"
                    >
                        註冊
                    </button>
                </div>
            </form>
        </div>
        <!--<router-link to="/regist"><span>註冊</span></router-link>-->
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
import { mapMutations } from 'vuex'

export default {
    name: 'Login',
    components: { Navbar },
    methods: {
        async login(e) {
            const name = e.target[0].value
            const password = e.target[1].value
            e.target[0].value = ''
            e.target[1].value = ''
            var res = await this.$store.dispatch('Login', { name, password })
            if (res) this.$router.push('/home')
        },
        async regist(e) {
            const name = e.target[0].value
            const password = e.target[1].value
            const password_ = e.target[2].value
            if (password == password_) {
                if(await this.$store.dispatch('NewUserAndLogin', { name, password }))this.$router.push('/home')
            } else this.alert('密碼不同')
        },
        ...mapMutations(['addAlert']),
        alert(msg) {
            this.$data.alertMsg = msg
        },
    }, //in order to avoid misspelling, I put blank here
    computed: {},
    data() {
        return {
            tab: 0,
            alertMsg: undefined,
        }
    },
}
</script>

<style></style>
