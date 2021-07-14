<template>
    <div>
        <navbar />
        <div :class="!alertMsg ? 'hidden' : ''">{{ alertMsg }}</div>
        <form v-on:submit.prevent="regist">
            <div class="grid grid-cols-2 border-blue">
                <label class="p-1 border-t-4 border-fuchsia-600 leading-10"
                    >帳戶</label
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
                <label class="p-1 border-t-4 border-fuchsia-600 leading-10"
                    >密碼</label
                ><input
                    type="text"
                    class="
                        focus:outline-none
                        p-1
                        border-t-4 border-fuchsia-600
                        leading-10
                    "
                    placeholder="輸入密碼"
                />
                <label
                    class="
                        p-1
                        border-t-4 border-b-4 border-fuchsia-600
                        leading-10
                    "
                    >確認密碼</label
                ><input
                    type="text"
                    class="
                        focus:outline-none
                        p-1
                        border-t-4 border-b-4 border-fuchsia-600
                        leading-10
                    "
                    placeholder="再次輸入密碼"
                />
            </div>
            <div>
                <input
                    type="submit"
                    class="bg-gray-100 hover:bg-gray-300 mt-4"
                />
            </div>
        </form>
    </div>
</template>

<script>
import Navbar from '@/components/navBar.vue'
// import { api } from '@/utils/api.js'

export default {
    name: 'Login',
    components: { Navbar },
    data: function () {
        return { alertMsg: undefined }
    },
    methods: {
        async regist(e) {
            const name = e.target[0].value
            const password = e.target[1].value
            const password_ = e.target[2].value
            if (password == password_) {
                await new this.$store.dispatch('NewUser', { name, password })
                this.$router.push('/login')
            } else this.alert('密碼無效')
        },
        alert(msg) {
            this.$data.alertMsg = msg
        },
    }, //in order to avoid misspelling, I put blank here
    computed: {},
}
</script>

<style>
form {
    padding: 10px;
}
.bor > label,
.bor > input {
    border: 1px solid black;
}
</style>
