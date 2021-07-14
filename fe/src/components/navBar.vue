<template>
    <div
        class="
            lg:sticky lg:h-auto lg:w-full
            h-full
            fixed
            top-0
            rounded-2xl
            z-30
        "
        id="app1"
    >
        <div class="absolute z-40 lg:hidden m-4">
            <!--lg+ hide-->
            <button @click="show = !show">
                <MenuIcon class="w-10 h-10" />
            </button>
        </div>
        <ul
            class="
                flex
                lg:w-full lg:h-auto lg:transform-none lg:p-2 lg:flex-row lg:ml-0
                flex-col
                p-10
                h-full
                justify-around
                text-2xl
                transition-all
                transform
                bg-gray-300
                w-60
            "
            :class="show ? 'ml-0' : '-ml-60'"
        >
            <form
                class="mx-2 my-2 p-2"
                @submit.prevent="query"
                v-show="!$route.path.match(/article\/find/)"
            >
                <input
                    class="focus:outline-none rounded-xl lg:w-auto w-32 px-4"
                    placeholder="尋找課程"
                />
            </form>
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2 p-2">
                <router-link to="/home" class="whitespace-nowrap"
                    >主頁</router-link
                >
            </li>
            <li
                class="hover:bg-blue-300 rounded-xl mx-2 my-2 p-2"
                v-show="!CheckTTL"
            >
                <router-link to="/login" class="whitespace-nowrap"
                    >登入</router-link
                >
            </li>
            <!--
            <li v-on:submit.prevent="logout" class="rounded-xl mx-2 my-2">
                <input
                    class="hover:bg-blue-300 whitespace-nowrap"
                    type="submit"
                />
            </li>--><!-- <li class="hover:bg-blue-300 rounded-xl mx-2 my-2">
                <router-link to="/regist" class="whitespace-nowrap"
                    >註冊</router-link
                >
            </li> -->
            <li class="hover:bg-blue-300 rounded-xl mx-2 my-2 p-2">
                <router-link to="/article/list" class="whitespace-nowrap"
                    >課程</router-link
                >
            </li>
            <li
                class="hover:bg-blue-300 rounded-xl mx-2 my-2 p-2"
                v-show="CheckTTL"
            >
                <router-link to="/account" class="whitespace-nowrap">
                {{user.name}}
                </router-link>
            </li>
        </ul>
    </div>
</template>
<script>
import { MenuIcon } from '@heroicons/vue/solid'
import { mapState, mapGetters } from 'vuex'
// import { api } from '@/utils/api.js'

export default {
    name: 'navBar',
    methods: {
        async query(e) {
            if (e.target[0].value.length >= 1)
                this.$router.push('/article/find/' + e.target[0].value)
        },
    },
    mounted() {
    },
    data() {
        return { show: false, accountName: '未登入'  }
    },
    computed: { ...mapGetters(['CheckTTL']) ,...mapState(['user'])},
    components: {
        MenuIcon,
    },
}
// function app1
// w-full
</script>
